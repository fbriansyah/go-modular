package userPort

import (
	"context"

	userModel "github.com/fbriansyah/go-modular/internal/model/user"
)

type UserService interface {
	CreateUser(ctx context.Context, req *userModel.CreateUserRequest) (*userModel.User, error)
	ListUser(ctx context.Context, query *userModel.ListUserQuery) ([]*userModel.User, error)
}
