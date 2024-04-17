package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	pb "github.com/MetaGigachad/task-tracker/tasks_service/internal/proto"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type server struct {
	pb.UnimplementedTasksServiceServer
	dbConn *pgxpool.Pool
}

func MakeErrorResponse(id int32, msg string) (*pb.TaskResponse, error) {
	return &pb.TaskResponse{
		Response: &pb.TaskResponse_Error{
			Error: &pb.Error{
				Code:    id,
				Message: msg,
			},
		},
	}, nil
}

func (s *server) CreateTask(ctx context.Context, req *pb.CreateTaskRequest) (*pb.TaskResponse, error) {
	log.Printf("Handling CreateTask")
	var id string
	createdAt := time.Now()
	status := pb.TaskStatus_Open
	err := s.dbConn.QueryRow(context.Background(), `
		INSERT INTO 
			tasks (creator_id, created_at, title, description, status)
		VALUES
			($1, $2, $3, $4, $5)
		RETURNING id`,
		req.UserId, createdAt, req.Title, req.Description, status).Scan(&id)
	if err != nil {
		return MakeErrorResponse(1, fmt.Sprintf("Database error: %v", err))
	}
	return &pb.TaskResponse{
		Response: &pb.TaskResponse_Task{
			Task: &pb.Task{
				Id:          id,
				CreatedAt:   timestamppb.New(createdAt),
				Title:       req.Title,
				Description: req.Description,
				Status:      status,
			},
		},
	}, nil

}

func (s *server) GetTask(ctx context.Context, req *pb.GetTaskRequest) (*pb.TaskResponse, error) {
	log.Printf("Handling GetTask")
	var (
		creatorId   string
		createdAt   time.Time
		title       string
		description string
		status      int32
	)
	err := s.dbConn.QueryRow(context.Background(), `
		SELECT
			creator_id, created_at, title, description, status
		FROM
			tasks
		WHERE
			id=$1`,
		req.TaskId).Scan(&creatorId, &createdAt, &title, &description, &status)
	if err != nil {
		return MakeErrorResponse(1, fmt.Sprintf("Database error: %v", err))
	}
	if creatorId != req.UserId {
		return MakeErrorResponse(2, "User is not a creator of this task")
	}
	return &pb.TaskResponse{
		Response: &pb.TaskResponse_Task{
			Task: &pb.Task{
				Id:          req.TaskId,
				CreatedAt:   timestamppb.New(createdAt),
				Title:       title,
				Description: description,
				Status:      pb.TaskStatus(status),
			},
		},
	}, nil
}

func (s *server) UpdateTask(ctx context.Context, req *pb.UpdateTaskRequest) (*pb.TaskResponse, error) {
	log.Printf("Handling UpdateTask")
	var (
		createdAt   time.Time
		title       string
		description string
		status      int32
	)
	err := s.dbConn.QueryRow(context.Background(), `
		UPDATE
			tasks
		SET
			title=COALESCE($1, title),
			description=COALESCE($2, description)
		WHERE
			id=$3 AND creator_id=$4
		RETURNING
			created_at, title, description, status`,
		req.NewTitle, req.NewDescription, req.TaskId, req.UserId).Scan(&createdAt, &title, &description, &status)
	if err != nil {
		return MakeErrorResponse(1, fmt.Sprintf("Database error: %v", err))
	}
	return &pb.TaskResponse{
		Response: &pb.TaskResponse_Task{
			Task: &pb.Task{
				Id:          req.TaskId,
				CreatedAt:   timestamppb.New(createdAt),
				Title:       title,
				Description: description,
				Status:      pb.TaskStatus(status),
			},
		},
	}, nil
}

func (s *server) DeleteTask(ctx context.Context, req *pb.DeleteTaskRequest) (*pb.TaskResponse, error) {
	log.Printf("Handling DeleteTask")
	var (
		createdAt   time.Time
		title       string
		description string
		status      int32
	)
	err := s.dbConn.QueryRow(context.Background(), `
		DELETE FROM
			tasks
		WHERE
			id=$1 AND creator_id=$2
		RETURNING
			created_at, title, description, status`,
		req.TaskId, req.UserId).Scan(&createdAt, &title, &description, &status)
	if err != nil {
		return MakeErrorResponse(1, fmt.Sprintf("Database error: %v", err))
	}
	return &pb.TaskResponse{
		Response: &pb.TaskResponse_Task{
			Task: &pb.Task{
				Id:          req.TaskId,
				CreatedAt:   timestamppb.New(createdAt),
				Title:       title,
				Description: description,
				Status:      pb.TaskStatus(status),
			},
		},
	}, nil
}

func (s *server) GetTaskPage(ctx context.Context, req *pb.GetTaskPageRequest) (*pb.TaskPageResponse, error) {
	log.Printf("Handling GetTaskPage")
	result, err := s.dbConn.Query(context.Background(), `
		SELECT
			id, created_at, title, description, status
		FROM
			tasks
		WHERE
			creator_id=$1
		ORDER BY
			created_at
		LIMIT $2 OFFSET $3`,
		req.UserId, req.PageSize, req.StartId)
	if err != nil {
		return &pb.TaskPageResponse{
			Response: &pb.TaskPageResponse_Error{
				Error: &pb.Error{
					Code:    1,
					Message: fmt.Sprintf("Database error: %v", err),
				},
			},
		}, nil
	}
	taskPage := pb.TaskPage{}
	for result.Next() {
		task := pb.Task{}
		var createdAt time.Time
		err = result.Scan(&task.Id, &createdAt, &task.Title, &task.Description, &task.Status)
		if err != nil {
			return &pb.TaskPageResponse{
				Response: &pb.TaskPageResponse_Error{
					Error: &pb.Error{
						Code:    2,
						Message: fmt.Sprintf("Database error: %v", err),
					},
				},
			}, nil
		}
		task.CreatedAt = timestamppb.New(createdAt)
		taskPage.Tasks = append(taskPage.Tasks, &task)
	}
	return &pb.TaskPageResponse{
		Response: &pb.TaskPageResponse_TaskPage{
			TaskPage: &taskPage,
		},
	}, nil
}

func main() {
	var (
		host     = flag.String("host", "", "gRPC host")
		port     = flag.Int("port", 50051, "gRPC port to listen on")
		dbConfig = flag.String("db-config", "", "gRPC port to listen on")
	)
	flag.Parse()

	dbConn, err := pgxpool.New(context.Background(), *dbConfig)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer dbConn.Close()

	listener, err := net.Listen("tcp", fmt.Sprintf("%v:%d", *host, *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	srv := grpc.NewServer()
	pb.RegisterTasksServiceServer(srv, &server{dbConn: dbConn})
	reflection.Register(srv)

	log.Printf("server listening at with reflection %v", listener.Addr())
	if err := srv.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
