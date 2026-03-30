package users_service

import (
	"context"
	"fmt"

	"github.com/Mihail4531/golang-todo/internal/core/domain"
)

func (s *UserService) GetUser(ctx context.Context, userID int) (*domain.User, error) {

	userDomain, err := s.usersRepository.GetUser(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("get user from repository %w ", err)
	}
	return userDomain, nil
}
