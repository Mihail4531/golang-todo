package stat_transport_http

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Mihail4531/golang-todo/internal/core/domain"
	core_logger "github.com/Mihail4531/golang-todo/internal/core/logger"
	core_http_request "github.com/Mihail4531/golang-todo/internal/core/transport/http/request"
	core_http_response "github.com/Mihail4531/golang-todo/internal/core/transport/http/response"
)

type GetStatisticsResponse struct {
	TasksCreated               int
	TasksCompleted             int
	TasksCompletedRate         *float64
	TasksAverageCompletionTime *string
}

func (h *StatHTTPHandler) GetStat(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)
	userID, from, to, err := getUserFromToQueryParams(r)
	if err != nil {
		responseHandler.ErrorResponse(err, "invalid query params")
		return
	}
	stat, err := h.statServ.GetStat(ctx, userID, from, to)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get statistics")
		return
	}
	response := toDTOFromDomain(stat)
	responseHandler.JsonResponse(response, http.StatusOK)

}
func toDTOFromDomain(statistics domain.Stat) GetStatisticsResponse {
	var avgTime *string
	if statistics.TasksAverageCompletionTime != nil {
		duration := statistics.TasksAverageCompletionTime.String()
		avgTime = &duration
	}

	return GetStatisticsResponse{
		TasksCreated:               statistics.TasksCreated,
		TasksCompleted:             statistics.TasksCompleted,
		TasksCompletedRate:         statistics.TasksCompletedRate,
		TasksAverageCompletionTime: avgTime,
	}
}
func getUserFromToQueryParams(r *http.Request) (*int, *time.Time, *time.Time, error) {
	userID, err := core_http_request.GetIntQueryParam(r, "user_id")
	if err != nil {
		return nil, nil, nil, fmt.Errorf("invalid query params user_id: %w", err)
	}
	fromTime, err := core_http_request.GetDateQueryParam(r, "from")
	if err != nil {
		return nil, nil, nil, fmt.Errorf("invalid query params from: %w", err)
	}
	toTime, err := core_http_request.GetDateQueryParam(r, "to")
	if err != nil {
		return nil, nil, nil, fmt.Errorf("invalid query params to: %w", err)
	}
	return userID, fromTime, toTime, nil
}
