package userService

import (
	"context"

	userModel "github.com/fbriansyah/go-modular/internal/model/user"
	"github.com/fbriansyah/go-modular/utils"
)

func (s *UserService) CreateUser(ctx context.Context, req *userModel.CreateUserRequest) (*userModel.User, error) {
	uuid := utils.GenerateUUID()
	user, err := userModel.NewUser(uuid, req.Email, req.Password, req.FirstName, req.LastName)
	if err != nil {
		return nil, err
	}

	err = s.userRepository.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
