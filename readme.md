# go-api-template

## Microservice

Simple microservice created to be a sample.
This service runs both an HTTP API server and a RabbitMQ consumer concurrently in the same process.

### Features
- RESTful API for task management
- MongoDB database integration
- RabbitMQ message consumer for async task processing
- Graceful shutdown handling

## Prerequisites

- Go 1.x or higher
- MongoDB instance
- RabbitMQ instance

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
```

3. Run the service:
```bash
make run
```

The service will start both the HTTP API server and RabbitMQ consumer simultaneously.

## API Examples

### Create a Task (HTTP)
```bash
curl -X POST -d '{"ID": "1", "Title": "New task", "Done": false}' -H "Content-Type: application/json" localhost:8080/tasks
```

### Get All Tasks
```bash
curl localhost:8080/tasks
```

### Delete a Task
```bash
curl -X DELETE "localhost:8080/tasks?id=1"
```

## RabbitMQ Consumer

The consumer listens to the configured RabbitMQ queue and automatically processes incoming task messages. Tasks received through RabbitMQ are created in the database.