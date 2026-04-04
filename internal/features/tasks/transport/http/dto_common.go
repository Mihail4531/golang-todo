package tasks_transport_http

import (
	"time"

	"github.com/Mihail4531/golang-todo/internal/core/domain"
)

type TaskDTOResposne struct {
	ID           int        `json:"id"`
	Version      int        `json:"version"`
	Title        string     `json:"title"`
	Description  *string    `json:"description"`
	Completed    bool       `json:"completed"`
	CreatedAt    time.Time  `json:"createdAt"`
	CompletedAt  *time.Time `json:"completedAt"`
	AuthorUserId int        `json:"author_id"`
}

func taskDTOFromDomain(taskDomain domain.Task) TaskDTOResposne {
	return TaskDTOResposne{
		ID:           taskDomain.ID,
		Version:      taskDomain.Version,
		Title:        taskDomain.Title,
		Description:  taskDomain.Description,
		Completed:    taskDomain.Completed,
		CreatedAt:    taskDomain.CreatedAt,
		CompletedAt:  taskDomain.CompletedAt,
		AuthorUserId: taskDomain.AuthorUserId,
	}
}

func tasksDTOFromDomain(tasksDomain []domain.Task) []TaskDTOResposne {
	tasksDto := make([]TaskDTOResposne, len(tasksDomain))
	for i, task := range tasksDomain {
		tasksDto[i] = taskDTOFromDomain(task)
	}
	return tasksDto
}
