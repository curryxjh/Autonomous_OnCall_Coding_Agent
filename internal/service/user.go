package service

import (
	"Autonomous_OnCall_Coding_Agent/internal/domain/user"
	"context"
)

type UserService interface {
	SignUp(ctx context.Context, u user.User) (string, error)
	Login(ctx context.Context, email, password string) (user.User, error)
	FindOrCreate(ctx context.Context, u user.User) (user.User, error)
	UpdateByID(ctx context.Context, u user.User) error
}

type userService struct {
}
