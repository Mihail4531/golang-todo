package tasks_transport_http

import (
	"net/http"

	core_logger "github.com/Mihail4531/golang-todo/internal/core/logger"
	core_http_request "github.com/Mihail4531/golang-todo/internal/core/transport/http/request"
	core_http_response "github.com/Mihail4531/golang-todo/internal/core/transport/http/response"
)

type GetTaskResponse TaskDTOResposne

func (h *TasksHTTPHandler) GetTask(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)
	taskID, err := core_http_request.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "invalid path value")
		return
	}
	task, err := h.tasksService.GetTask(ctx, taskID)
	if err != nil {
		responseHandler.ErrorResponse(err, "get task")
		return
	}
	response := GetTaskResponse(taskDTOFromDomain(task))
	responseHandler.JsonResponse(response, http.StatusOK)
}
