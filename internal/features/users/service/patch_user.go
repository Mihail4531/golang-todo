package users_service

import (
	"context"
	"fmt"

	"github.com/Mihail4531/golang-todo/internal/core/domain"
)

func (s *UserService) PatchUser(ctx context.Context, userID int, patch *domain.UserPatch) (*domain.User, error) {
	user, err := s.usersRepository.GetUser(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("get user %w", err)
	}
	if err := user.ApplayPatch(patch); err != nil {
		return nil, fmt.Errorf("applay patch user")
	}
	userPatch, err := s.usersRepository.PatchUser(ctx, userID, user)
	if err != nil {
		return nil, fmt.Errorf(" patch user %w", err)
	}
	return userPatch, nil
}
