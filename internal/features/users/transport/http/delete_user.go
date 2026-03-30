package users_transport_http

import (
	"net/http"

	core_logger "github.com/Mihail4531/golang-todo/internal/core/logger"
	core_http_request "github.com/Mihail4531/golang-todo/internal/core/transport/http/request"
	core_http_response "github.com/Mihail4531/golang-todo/internal/core/transport/http/response"
)

func (h *UsersHTTPHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHander := core_http_response.NewHTTPResponseHandler(log, w)
	userID, err := core_http_request.GetIntPathValue(r, "id")
	if err != nil {
		responseHander.ErrorResponse(err, "failed to get useId to path value")
		return
	}
	err = h.usersSerice.DeleteUser(ctx, userID)
	if err != nil {
		responseHander.ErrorResponse(err, "failed to delete user")
	}
	responseHander.NoContentResponse()
}
