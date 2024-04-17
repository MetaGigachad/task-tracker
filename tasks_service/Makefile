list-services:
	grpcurl -import-path ./api -proto tasks_service.proto list
list-methods:
	grpcurl -import-path ./api -proto tasks_service.proto list tasks_service.TasksService
create-task:
	grpcurl -import-path ./api -proto tasks_service.proto -plaintext -d '{"user_id": 0, "title": "Task1", "description": "desc"}' \
		0.0.0.0:50051 tasks_service.TasksService/CreateTask
