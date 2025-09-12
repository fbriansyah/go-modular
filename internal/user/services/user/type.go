package userService

import (
	"github.com/fbriansyah/go-modular/config"
	userPort "github.com/fbriansyah/go-modular/ports/user"
)

type UserService struct {
	conf           *config.Config
	userRepository userPort.UserRepository
}

type Option func(*UserService)

func NewUserService(conf *config.Config, opts ...Option) *UserService {
	userService := &UserService{conf: conf}
	for _, opt := range opts {
		opt(userService)
	}
	return userService
}

var _ userPort.UserService = (*UserService)(nil)
