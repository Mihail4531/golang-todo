package tasks_transport_http

import (
	"context"
	"net/http"

	"github.com/Mihail4531/golang-todo/internal/core/domain"
	core_http_server "github.com/Mihail4531/golang-todo/internal/core/transport/server"
)

type TasksService interface {
	CreateTask(ctx context.Context, task domain.Task) (domain.Task, error)
	GetTasks(ctx context.Context, limit *int, offset *int, userID *int) ([]domain.Task, error)
	GetTask(ctx context.Context, taskID int) (domain.Task, error)
	DeleteTask(ctx context.Context, taskID int) error
	PatchTask(ctx context.Context, patch domain.TaskPatch, taskID int) (domain.Task, error)
}

type TasksHTTPHandler struct {
	tasksService TasksService
}

func NewTasksHTTPHandler(tasksService TasksService) *TasksHTTPHandler {
	return &TasksHTTPHandler{
		tasksService: tasksService,
	}
}
func (h *TasksHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		core_http_server.NewRoute(http.MethodPost, "/tasks", h.CreateTasks),
		core_http_server.NewRoute(http.MethodGet, "/tasks", h.GetTasks),
		core_http_server.NewRoute(http.MethodGet, "/tasks/{id}", h.GetTask),
		core_http_server.NewRoute(http.MethodDelete, "/tasks/{id}", h.DeleteTask),
		core_http_server.NewRoute(http.MethodPatch, "/task/{id}", h.PatchTasks),
	}
}
