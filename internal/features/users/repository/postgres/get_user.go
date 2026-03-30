package users_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/Mihail4531/golang-todo/internal/core/domain"
	core_errors "github.com/Mihail4531/golang-todo/internal/core/errors"
	core_postgres_pool "github.com/Mihail4531/golang-todo/internal/core/repository/postgres/pool"
)

func (r *UsersRepository) GetUser(ctx context.Context, userID int) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()
	query := `SELECT id, version, full_name, phone_number  FROM todoapp.users WHERE id = $1 `
	var userModel UserModel
	err := r.pool.QueryRow(ctx, query, userID).Scan(&userModel.ID, &userModel.Version, &userModel.FullName, &userModel.PhoneNumber)
	if err != nil {
		if errors.Is(err, core_postgres_pool.ErrNoRows) {
			return nil, fmt.Errorf("user with id='%d' %w", userID, core_errors.ErrNotFound)
		}
		return nil, fmt.Errorf("scan error: %w", err)
	}
	userDomain := domain.NewUser(userModel.ID, userModel.Version, userModel.FullName, userModel.PhoneNumber)
	return userDomain, nil
}
