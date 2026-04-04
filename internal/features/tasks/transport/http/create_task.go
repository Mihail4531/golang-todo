package tasks_transport_http

import (
	"net/http"

	"github.com/Mihail4531/golang-todo/internal/core/domain"
	core_logger "github.com/Mihail4531/golang-todo/internal/core/logger"
	core_http_request "github.com/Mihail4531/golang-todo/internal/core/transport/http/request"
	core_http_response "github.com/Mihail4531/golang-todo/internal/core/transport/http/response"
)

type CreateTaskRequest struct {
	Title        string  `json:"title" validate:"required,min=1,max=100"`
	Description  *string `json:"description" validate:"omitempty,min=1,max=1000"`
	AuthorUserId int     `json:"author_id" validate:"required"`
}
type CreateTaskResponse TaskDTOResposne

func (h *TasksHTTPHandler) CreateTasks(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)
	var request CreateTaskRequest
	if err := core_http_request.DecodeValid(r, &request); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode and validate http request")
		return
	}
	taskDomain := domain.NewTaskUninitialized(request.Title, request.Description, request.AuthorUserId)
	task, err := h.tasksService.CreateTask(ctx, taskDomain)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to create task")
		return
	}
	response := CreateTaskResponse(taskDTOFromDomain(task))
	responseHandler.JsonResponse(response, http.StatusCreated)
}
