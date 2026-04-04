package tasks_postgres_repository

import (
	"context"
	"fmt"

	"github.com/Mihail4531/golang-todo/internal/core/domain"
)

func (r *TasksRepository) GetTasks(ctx context.Context, limit *int, offset *int, userID *int) ([]domain.Task, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()
	query := `SELECT id, version, title, description, completed, created_at, completed_at, author_user_id 
	FROM todoapp.tasks 
	%s
	ORDER BY id ASC 
	LIMIT $1
	OFFSET $2`
	args := []any{limit, offset}
	if userID != nil {
		query = fmt.Sprintf(query, "WHERE author_user_id=$3")
		args = append(args, userID)
	} else {
		query = fmt.Sprintf(query, "")
	}
	rows, err := r.pool.Query(ctx, query, args...)
	if err != nil {
		return []domain.Task{}, fmt.Errorf("select users: %w", err)
	}
	defer rows.Close()
	var tasksModels []TaskModel
	for rows.Next() {
		var taskModel TaskModel
		if err := rows.Scan(&taskModel.ID, &taskModel.Version, &taskModel.Title, &taskModel.Description, &taskModel.Completed, &taskModel.CreatedAt, &taskModel.CompletedAt, &taskModel.AuthorUserId); err != nil {

			return []domain.Task{}, fmt.Errorf("scan error: %w", err)
		}
		tasksModels = append(tasksModels, taskModel)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("next rows %w", err)
	}
	taskDomain := tasksModelFromDomains(tasksModels)
	return taskDomain, nil
}
