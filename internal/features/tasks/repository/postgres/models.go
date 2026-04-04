package tasks_postgres_repository

import (
	"time"

	"github.com/Mihail4531/golang-todo/internal/core/domain"
)

type TaskModel struct {
	ID           int
	Version      int
	Title        string
	Description  *string
	Completed    bool
	CreatedAt    time.Time
	CompletedAt  *time.Time
	AuthorUserId int
}

func tasksModelFromDomains(tasks []TaskModel) []domain.Task {
	taskDomain := make([]domain.Task, len(tasks))
	for i, task := range tasks {
		taskDomain[i] = domain.NewTask(task.ID, task.Version, task.Title, task.Description, task.Completed, task.CreatedAt, task.CompletedAt, task.AuthorUserId)
	}
	return taskDomain
}
func taskModelFromDomain(task TaskModel) domain.Task {
	return domain.NewTask(
		task.ID,
		task.Version,
		task.Title,
		task.Description,
		task.Completed,
		task.CreatedAt,
		task.CompletedAt,
		task.AuthorUserId)
}
