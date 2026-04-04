package tasks_transport_http

import (
	"fmt"
	"net/http"

	"github.com/Mihail4531/golang-todo/internal/core/domain"
	core_logger "github.com/Mihail4531/golang-todo/internal/core/logger"
	core_http_request "github.com/Mihail4531/golang-todo/internal/core/transport/http/request"
	core_http_response "github.com/Mihail4531/golang-todo/internal/core/transport/http/response"
	core_http_types "github.com/Mihail4531/golang-todo/internal/core/transport/http/types"
)

type PatchTaskRequest struct {
	Title       core_http_types.Nullable[string] `json:"title"`
	Description core_http_types.Nullable[string] `json:"description"`
	Completed   core_http_types.Nullable[bool]   `json:"completed"`
}
type PatchTaskResponse TaskDTOResposne

func (p *PatchTaskRequest) Validate() error {
	if p.Title.Set {
		if p.Title.Value == nil {
			return fmt.Errorf("title required")
		}
		titleLen := len([]rune(*p.Title.Value))
		if titleLen < 1 || titleLen > 100 {
			return fmt.Errorf("Title must be between 3 and 100 symbols")
		}
	}
	if p.Description.Set {
		if p.Description.Value != nil {
			descriptionLen := len([]rune(*p.Description.Value))
			if descriptionLen < 1 || descriptionLen > 100 {
				return fmt.Errorf("Description must be between 1 and 1000 symbols")
			}
		}
	}
	if p.Completed.Set {
		if p.Completed.Value == nil {
			return fmt.Errorf("completed is required")
		}
	}
	return nil
}
func (h *TasksHTTPHandler) PatchTasks(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)
	taskID, err := core_http_request.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "invalid path value")
		return
	}
	var request PatchTaskRequest
	if err := core_http_request.DecodeValid(r, &request); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode request")
		return
	}
	taskPatch := taskPatchFromRequest(request)
	task, err := h.tasksService.PatchTask(ctx, taskPatch, taskID)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to patch")
		return
	}
	response := PatchTaskResponse(taskDTOFromDomain(task))
	responseHandler.JsonResponse(response, http.StatusOK)
}
func taskPatchFromRequest(request PatchTaskRequest) domain.TaskPatch {
	return domain.NewTaskPatch(request.Title.Nullable, request.Description.Nullable, request.Completed.Nullable)
}
