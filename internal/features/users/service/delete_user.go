package users_service

import (
	"context"
	"fmt"
)

func (s *UserService) DeleteUser(ctx context.Context, userID int) error {
	err := s.usersRepository.DeleteUser(ctx, userID)
	if err != nil {
		return fmt.Errorf("delete user  %w ", err)
	}
	return nil
}
