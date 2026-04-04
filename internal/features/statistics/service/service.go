package stat_service

import (
	"context"
	"time"

	"github.com/Mihail4531/golang-todo/internal/core/domain"
)

type StatRepository interface {
	GetTasks(ctx context.Context, userID *int, from *time.Time, to *time.Time) ([]domain.Task, error)
}
type StatService struct {
	statRepository StatRepository
}

func NewStatService(statRepository StatRepository) *StatService {
	return &StatService{
		statRepository: statRepository,
	}
}
