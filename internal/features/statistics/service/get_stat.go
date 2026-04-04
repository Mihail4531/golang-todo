package stat_service

import (
	"context"
	"fmt"
	"time"

	"github.com/Mihail4531/golang-todo/internal/core/domain"
	core_errors "github.com/Mihail4531/golang-todo/internal/core/errors"
)

func (s *StatService) GetStat(ctx context.Context, userID *int, from *time.Time, to *time.Time) (domain.Stat, error) {
	if from != nil && to != nil {
		if to.Before(*from) || to.Equal(*from) {
			return domain.Stat{}, fmt.Errorf("to must be after from : %w", core_errors.ErrInvalidArgument)
		}
	}
	tasks, err := s.statRepository.GetTasks(ctx, userID, from, to)
	if err != nil {
		return domain.Stat{}, fmt.Errorf("get stats: %w", err)
	}
	stat := calcStatistics(tasks)
	return stat, nil

}

func calcStatistics(tasks []domain.Task) domain.Stat {
	if len(tasks) == 0 {
		return domain.Stat{
			TasksCreated:               0,
			TasksCompleted:             0,
			TasksCompletedRate:         nil,
			TasksAverageCompletionTime: nil,
		}
	}
	tasksCreated := len(tasks)
	tasksCompleted := 0
	var totalCompletedDuration time.Duration
	for _, task := range tasks {
		if task.Completed {
			tasksCompleted++

		}
		completionDuration := task.CompletedDuration()
		if completionDuration != nil {
			totalCompletedDuration += *completionDuration
		}
	}
	tasksCompletedRate := float64(tasksCompleted) / float64(tasksCreated) * 100
	var taskAverageCompletionTime *time.Duration
	if tasksCompleted > 0 && totalCompletedDuration != 0 {
		avg := totalCompletedDuration / time.Duration(tasksCompleted)
		taskAverageCompletionTime = &avg
	}
	return domain.Stat{
		TasksCreated:               tasksCreated,
		TasksCompleted:             tasksCompleted,
		TasksCompletedRate:         &tasksCompletedRate,
		TasksAverageCompletionTime: taskAverageCompletionTime,
	}
}
