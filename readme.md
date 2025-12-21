# go-api-template

## Microservice

Simple microservice created to be a sample.
This service runs an HTTP API server, a gRPC server, and a RabbitMQ consumer concurrently in the same process.

### Features
- RESTful API for task management
- gRPC API for task management
- MongoDB database integration
- RabbitMQ message consumer for async task processing
- Graceful shutdown handling

## Prerequisites

- Go 1.x or higher
- MongoDB instance
- RabbitMQ instance
- Protocol Buffers compiler (protoc) - for regenerating gRPC code

## How to run?

1. Copy the example environment file and configure it:
```bash
cp env.example .env
```

2. Update the `.env` file with your configuration:
```
MONGO_URI=mongodb://localhost:27017/your-database-name
RABBITMQ_URI=amqp://guest:guest@localhost:5672/
RABBITMQ_EXCHANGE=TASKS_EXCHANGE
RABBITMQ_EXCHANGE_TYPE=topic
RABBITMQ_QUEUE=TASKS_QUEUE
RABBITMQ_ROUTING_KEY=task.create
PORT=8080
GRPC_PORT=50051
```

3. Run the service:
```bash
make run
```

The service will start the HTTP API server (port 8080), gRPC server (port 50051), and RabbitMQ consumer simultaneously.

## API Examples

### REST API (HTTP)

#### Create a Task
```bash
curl -X POST -d '{"ID": "1", "Title": "New task", "Done": false}' -H "Content-Type: application/json" localhost:8080/tasks
```

#### Get All Tasks
```bash
curl localhost:8080/tasks
```

#### Delete a Task
```bash
curl -X DELETE "localhost:8080/tasks?id=1"
```

### gRPC API

The gRPC service provides the same task management functionality through gRPC endpoints.

#### Using the Example Client
```bash
go run examples/grpc_client.go
```

#### gRPC Methods Available
- `GetTasks` - Retrieve all tasks
- `CreateTask` - Create a new task
- `UpdateTask` - Update an existing task
- `DeleteTask` - Delete a task by ID

#### Using grpcurl
If you have `grpcurl` installed, you can test the gRPC API:

```bash
grpcurl -plaintext localhost:50051 list

grpcurl -plaintext localhost:50051 task.TaskService/GetTasks

grpcurl -plaintext -d '{"title": "New gRPC task", "done": false}' localhost:50051 task.TaskService/CreateTask
```

## Protocol Buffers

The gRPC service is defined in `proto/task.proto`. If you modify the proto file, regenerate the Go code:

```bash
protoc --go_out=. --go_opt=paths=source_relative \
  --go-grpc_out=. --go-grpc_opt=paths=source_relative \
  proto/task.proto
```

## RabbitMQ Consumer

The consumer listens to the configured RabbitMQ queue and automatically processes incoming task messages. Tasks received through RabbitMQ are created in the database.