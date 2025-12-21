package grpc

import (
	"context"
	"log"
	"net"

	"github.com/matheuscaet/go-api-template/business"
	task "github.com/matheuscaet/go-api-template/business/types"
	"github.com/matheuscaet/go-api-template/internal/config"
	pb "github.com/matheuscaet/go-api-template/proto"
	"google.golang.org/grpc"
)

type taskServer struct {
	pb.UnimplementedTaskServiceServer
	service business.TaskService
}

func NewTaskServer() *taskServer {
	return &taskServer{
		service: business.NewTaskService(),
	}
}

func (s *taskServer) GetTasks(ctx context.Context, req *pb.GetTasksRequest) (*pb.GetTasksResponse, error) {
	tasks, err := s.service.GetTasks(ctx)
	if err != nil {
		return nil, err
	}

	log.Printf("Hit GetTasks gRPC method")

	pbTasks := make([]*pb.Task, len(tasks))
	for i, t := range tasks {
		pbTasks[i] = &pb.Task{
			Id:    t.ID,
			Title: t.Title,
			Done:  t.Done,
		}
	}

	return &pb.GetTasksResponse{
		Tasks: pbTasks,
	}, nil
}

func (s *taskServer) CreateTask(ctx context.Context, req *pb.CreateTaskRequest) (*pb.CreateTaskResponse, error) {
	newTask := task.Task{
		Title: req.Title,
		Done:  req.Done,
	}

	createdTask, err := s.service.CreateTask(ctx, newTask)
	if err != nil {
		return nil, err
	}

	return &pb.CreateTaskResponse{
		Task: &pb.Task{
			Id:    createdTask.ID,
			Title: createdTask.Title,
			Done:  createdTask.Done,
		},
	}, nil
}

func (s *taskServer) UpdateTask(ctx context.Context, req *pb.UpdateTaskRequest) (*pb.UpdateTaskResponse, error) {
	updateTask := task.Task{
		ID:    req.Id,
		Title: req.Title,
		Done:  req.Done,
	}

	updatedTask, err := s.service.UpdateTask(ctx, updateTask)
	if err != nil {
		return nil, err
	}

	return &pb.UpdateTaskResponse{
		Task: &pb.Task{
			Id:    updatedTask.ID,
			Title: updatedTask.Title,
			Done:  updatedTask.Done,
		},
	}, nil
}

func (s *taskServer) DeleteTask(ctx context.Context, req *pb.DeleteTaskRequest) (*pb.DeleteTaskResponse, error) {
	err := s.service.DeleteTask(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.DeleteTaskResponse{
		Message: "Task deleted successfully",
	}, nil
}

func StartGRPCServer() {
	port := config.GRPCPort
	if port == "" {
		port = "50051"
	}

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", port, err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterTaskServiceServer(grpcServer, NewTaskServer())

	log.Printf("gRPC server listening on port %s", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC: %v", err)
	}
}
