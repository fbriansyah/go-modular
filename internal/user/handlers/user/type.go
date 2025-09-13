package userHandler

import (
	"github.com/gofiber/fiber/v2"

	"github.com/fbriansyah/go-modular/config"
	userService "github.com/fbriansyah/go-modular/internal/user/services/user"
)

type Option func(*UserHandler)

type UserHandler struct {
	httpApp     *fiber.App
	userService *userService.UserService
}

func NewUserHandler(conf *config.Config, opts ...Option) *UserHandler {
	userHandler := &UserHandler{}
	for _, opt := range opts {
		opt(userHandler)
	}
	return userHandler
}

func WithUserService(userService *userService.UserService) Option {
	return func(u *UserHandler) {
		u.userService = userService
	}
}
