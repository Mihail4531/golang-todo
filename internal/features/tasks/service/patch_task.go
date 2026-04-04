package tasks_service

import (
	"context"
	"fmt"

	"github.com/Mihail4531/golang-todo/internal/core/domain"
)

func (s *TasksService) PatchTask(ctx context.Context, patch domain.TaskPatch, taskID int) (domain.Task, error) {
	task, err := s.GetTask(ctx, taskID)
	if err != nil {
		return domain.Task{}, fmt.Errorf("get task from repository: %w", err)
	}
	if err := task.ApplyPatch(&patch); err != nil {
		return domain.Task{}, fmt.Errorf("apply patch: %w", err)
	}
	taskPatched, err := s.tasksRepository.PatchTask(ctx, task, taskID)
	if err != nil {
		return domain.Task{}, fmt.Errorf("patch task %w", err)
	}
	return taskPatched, nil
}
