package users_transport_http

import (
	"context"
	"net/http"

	"github.com/Mihail4531/golang-todo/internal/core/domain"
	core_http_server "github.com/Mihail4531/golang-todo/internal/core/transport/server"
)

type UsersService interface {
	CreateUser(ctx context.Context, user *domain.User) (*domain.User, error)
	GetUsers(ctx context.Context, limit *int, offset *int) ([]domain.User, error)
	GetUser(ctx context.Context, userID int) (*domain.User, error)
	DeleteUser(ctx context.Context, userID int) error
	PatchUser(ctx context.Context, userID int, patch *domain.UserPatch) (*domain.User, error)
}
type UsersHTTPHandler struct {
	usersSerice UsersService
}

func NewUsersHTTPHandler(usersService UsersService) *UsersHTTPHandler {
	return &UsersHTTPHandler{
		usersSerice: usersService,
	}
}
func (h *UsersHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		core_http_server.NewRoute(http.MethodPost, "/users", h.CreateUser),
		core_http_server.NewRoute(http.MethodGet, "/users", h.GetUsers),
		core_http_server.NewRoute(http.MethodGet, "/users/{id}", h.GetUser),
		core_http_server.NewRoute(http.MethodDelete, "/users/{id}", h.DeleteUser),
		core_http_server.NewRoute(http.MethodPatch, "/users/{id}", h.PatchUser),
	}
}
