package service

import (
	"Autonomous_OnCall_Coding_Agent/internal/domain/user"
	"context"
)

type UserService interface {
	SignUp(ctx context.Context, u user.User) error
	Login(ctx context.Context, email, password string) error
	FindOrCreate(ctx context.Context, u user.User) (user.User, error)
	UpdateByID(ctx context.Context, u user.User) error
}

type userService struct {
}
