package business

import (
	"context"

	models "github.com/matheuscaet/go-api-template/business/models"
	task "github.com/matheuscaet/go-api-template/business/types"
)

type TaskService interface {
	GetTasks(ctx context.Context) ([]task.Task, error)
	CreateTask(ctx context.Context, task task.Task) (task.Task, error)
	UpdateTask(ctx context.Context, task task.Task) (task.Task, error)
	DeleteTask(ctx context.Context, id string) error
}

type taskService struct {
}

func NewTaskService() TaskService {
	return &taskService{}
}

func (s *taskService) GetTasks(ctx context.Context) ([]task.Task, error) {
	return models.GetTasks(ctx)
}

func (s *taskService) CreateTask(ctx context.Context, task task.Task) (task.Task, error) {
	return models.CreateTask(ctx, task)
}

func (s *taskService) UpdateTask(ctx context.Context, task task.Task) (task.Task, error) {
	return models.UpdateTask(ctx, task)
}

func (s *taskService) DeleteTask(ctx context.Context, id string) error {
	return models.DeleteTask(ctx, id)
}
