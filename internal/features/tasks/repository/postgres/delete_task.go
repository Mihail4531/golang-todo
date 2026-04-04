package tasks_postgres_repository

import (
	"context"
	"fmt"

	core_errors "github.com/Mihail4531/golang-todo/internal/core/errors"
)

func (r *TasksRepository) DeleteTask(ctx context.Context, taskID int) error {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `DELETE FROM todoapp.tasks WHERE id = $1 `
	commandTag, err := r.pool.Exec(ctx, query, taskID)
	if err != nil {
		return fmt.Errorf("postgres error: %w", err)
	}
	if commandTag.RowsAffected() == 0 {
		return fmt.Errorf("task with id=%d not found: %w", taskID, core_errors.ErrNotFound)
	}
	return nil
}
