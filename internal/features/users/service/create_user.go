package users_service

import (
	"context"
	"fmt"

	"github.com/Mihail4531/golang-todo/internal/core/domain"
)

func (s *UserService) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	if err := user.Validate(); err != nil {
		return nil, fmt.Errorf("validate user domain:%w", err)
	}
	user, err := s.usersRepository.CreateUser(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("create user %w", err)
	}
	return user, nil
}
