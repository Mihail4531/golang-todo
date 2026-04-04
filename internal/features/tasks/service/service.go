package tasks_service

import (
	"context"

	"github.com/Mihail4531/golang-todo/internal/core/domain"
)

type TasksRepository interface {
	CreateTask(ctx context.Context, task domain.Task) (domain.Task, error)
	GetTasks(ctx context.Context, limit *int, offset *int, userID *int) ([]domain.Task, error)
	GetTask(ctx context.Context, taskID int) (domain.Task, error)
	DeleteTask(ctx context.Context, taskID int) error
	PatchTask(ctx context.Context, task domain.Task, taskID int) (domain.Task, error)
}
type TasksService struct {
	tasksRepository TasksRepository
}

func NewTasksService(tasksRepository TasksRepository) *TasksService {
	return &TasksService{
		tasksRepository: tasksRepository,
	}
}
