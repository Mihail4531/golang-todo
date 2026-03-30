package users_service

import (
	"context"

	"github.com/Mihail4531/golang-todo/internal/core/domain"
)

type UsersRepository interface {
	CreateUser(ctx context.Context, user *domain.User) (*domain.User, error)
	GetUsers(ctx context.Context, limit *int, offset *int) ([]domain.User, error)
	GetUser(ctx context.Context, userID int) (*domain.User, error)
	DeleteUser(ctx context.Context, userID int) error
	PatchUser(ctx context.Context, id int, user *domain.User) (*domain.User, error)
}
type UserService struct {
	usersRepository UsersRepository
}

func NewUsersService(usersRepository UsersRepository) *UserService {
	return &UserService{
		usersRepository: usersRepository,
	}
}
