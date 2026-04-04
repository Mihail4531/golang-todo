package stat_transport_http

import (
	"context"
	"net/http"
	"time"

	"github.com/Mihail4531/golang-todo/internal/core/domain"
	core_http_server "github.com/Mihail4531/golang-todo/internal/core/transport/server"
)

type StatService interface {
	GetStat(ctx context.Context, userID *int, from *time.Time, to *time.Time) (domain.Stat, error)
}

type StatHTTPHandler struct {
	statServ StatService
}

func NewStatHTTPHandler(statServ StatService) *StatHTTPHandler {
	return &StatHTTPHandler{
		statServ: statServ,
	}
}
func (h *StatHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		core_http_server.NewRoute(http.MethodGet, "/statistics", h.GetStat),
	}
}
