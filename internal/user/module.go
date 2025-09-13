package userModule

import (
	"github.com/fbriansyah/go-modular/config"
	sharedModule "github.com/fbriansyah/go-modular/internal/shared"
	userHandler "github.com/fbriansyah/go-modular/internal/user/handlers/user"
	userRepository "github.com/fbriansyah/go-modular/internal/user/repositories/user"
	userService "github.com/fbriansyah/go-modular/internal/user/services/user"
	"github.com/fbriansyah/go-modular/pkg/database"
	"github.com/gofiber/fiber/v2"
)

type UserModule struct {
	conf    *config.Config
	httpApp *fiber.App
	db      *database.DB
}

type Option func(*UserModule)

func WithDB(db *database.DB) Option {
	return func(u *UserModule) {
		u.db = db
	}
}

func WithHTTPApp(httpApp *fiber.App) Option {
	return func(u *UserModule) {
		u.httpApp = httpApp
	}
}

func NewUserModule(conf *config.Config, opts ...Option) *UserModule {
	userModule := &UserModule{
		conf: conf,
	}
	for _, opt := range opts {
		opt(userModule)
	}
	return userModule
}

func (um *UserModule) Run() {
	userRepo := userRepository.NewUserRepository(um.db)
	userService := userService.NewUserService(
		um.conf,
		userService.WithUserRepository(userRepo),
	)

	userHandler := userHandler.NewUserHandler(
		um.conf,
		userHandler.WithUserService(userService),
	)
	userHandler.SetupRoutes(um.httpApp)

}

var _ sharedModule.Application = (*UserModule)(nil)
