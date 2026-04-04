package domain

import "time"

type Stat struct {
	TasksCreated               int
	TasksCompleted             int
	TasksCompletedRate         *float64
	TasksAverageCompletionTime *time.Duration
}

func NewStat(tasksCreated int, tasksCompleted int, tasksCompletedRate *float64, tasksAverageCompletionTime *time.Duration) Stat {
	return Stat{
		TasksCreated:               tasksCreated,
		TasksCompleted:             tasksCompleted,
		TasksCompletedRate:         tasksCompletedRate,
		TasksAverageCompletionTime: tasksAverageCompletionTime,
	}
}
