package users_transport_http

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/Mihail4531/golang-todo/internal/core/domain"
	core_logger "github.com/Mihail4531/golang-todo/internal/core/logger"
	core_http_request "github.com/Mihail4531/golang-todo/internal/core/transport/http/request"
	core_http_response "github.com/Mihail4531/golang-todo/internal/core/transport/http/response"
	core_http_types "github.com/Mihail4531/golang-todo/internal/core/transport/http/types"
	core_http_utils "github.com/Mihail4531/golang-todo/internal/core/transport/http/utils"
)

type PatchUserRequest struct {
	FullName    core_http_types.Nullable[string] `json:"full_name"`
	PhoneNumber core_http_types.Nullable[string] `json:"phone_number"`
}

func (r *PatchUserRequest) Validate() error {
	if r.FullName.Set {
		if r.FullName.Value == nil {
			return fmt.Errorf("Full name cant be null")
		}
		fullNameLen := len([]rune(*r.FullName.Value))
		if fullNameLen < 3 || fullNameLen > 100 {
			return fmt.Errorf("FullName must be between 3 and 100 symbols")
		}
	}
	if r.PhoneNumber.Set {
		if r.PhoneNumber.Value != nil {
			phoneNumberLen := len([]rune(*r.PhoneNumber.Value))
			if phoneNumberLen < 10 || phoneNumberLen > 15 {
				return fmt.Errorf("PhoneNumber must be between 10 and 15 symbols")
			}
			if !strings.HasPrefix(*r.PhoneNumber.Value, "+") {
				return fmt.Errorf("PhoneNumber must startwith + symbols")
			}
		}
	}
	return nil
}

type PatchUserResponse UserDTOResponse

func (h *UsersHTTPHandler) PatchUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)
	userId, err := core_http_utils.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get useId to path value")
		return
	}
	var request PatchUserRequest
	if err := core_http_request.DecodeValid(r, &request); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode and validate http request")
		return
	}
	userPatch := userPatchFromRequest(request)
	userDomain, err := h.usersSerice.PatchUser(ctx, userId, userPatch)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed patch user")
		return
	}
	response := PatchUserResponse(userDTOFromDomain(userDomain))
	responseHandler.JsonResponse(response, http.StatusOK)
}
func userPatchFromRequest(request PatchUserRequest) *domain.UserPatch {
	return &domain.UserPatch{
		FullName:    request.FullName.ToDomain(),
		PhoneNumber: request.PhoneNumber.ToDomain(),
	}
}
