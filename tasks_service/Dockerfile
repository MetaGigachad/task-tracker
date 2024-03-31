FROM golang:1.22.1 as build

WORKDIR /app
COPY . .
ENV CGO_ENABLED=0
RUN go build ./cmd/tasks_service 

FROM scratch

WORKDIR /app
COPY --from=build /app/tasks_service .
ENTRYPOINT [ "./tasks_service" ]
