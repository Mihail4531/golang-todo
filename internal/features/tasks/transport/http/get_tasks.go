package tasks_transport_http

import (
	"fmt"
	"net/http"

	core_logger "github.com/Mihail4531/golang-todo/internal/core/logger"
	core_http_request "github.com/Mihail4531/golang-todo/internal/core/transport/http/request"
	core_http_response "github.com/Mihail4531/golang-todo/internal/core/transport/http/response"
)

type GetTasksResponse []TaskDTOResposne

func (h *TasksHTTPHandler) GetTasks(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)
	userID, limit, offset, err := getLimitOffsetQueryParams(r)
	if err != nil {
		responseHandler.ErrorResponse(err, "invalid query params")
		return
	}
	tasksDomain, err := h.tasksService.GetTasks(ctx, limit, offset, userID)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get users")
	}
	response := GetTasksResponse(tasksDTOFromDomain(tasksDomain))
	responseHandler.JsonResponse(response, http.StatusOK)
}
func getLimitOffsetQueryParams(r *http.Request) (*int, *int, *int, error) {
	limit, err := core_http_request.GetIntQueryParam(r, "limit")
	if err != nil {
		return nil, nil, nil, fmt.Errorf("get 'limit' param: %w", err)
	}
	userID, err := core_http_request.GetIntQueryParam(r, "UserId")
	if err != nil {
		return nil, nil, nil, fmt.Errorf("get'user_id' param: %w", err)
	}
	offset, err := core_http_request.GetIntQueryParam(r, "offset")
	if err != nil {
		return nil, nil, nil, fmt.Errorf("get 'offset' param: %w", err)
	}
	return userID, limit, offset, nil
}
