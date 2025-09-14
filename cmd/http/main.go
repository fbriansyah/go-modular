package http

import (
	"github.com/fbriansyah/go-modular/config"
	userModule "github.com/fbriansyah/go-modular/internal/user"
	"github.com/fbriansyah/go-modular/pkg/database"
	"github.com/gofiber/fiber/v2"
)

func Main() {
	conf := config.Config{}
	secret := config.Secret{}
	dbManager, err := database.NewManager(&conf.Database, &secret.Database, "./migrations")
	if err != nil {
		panic(err)
	}
	httpApp := fiber.New()
	userModel := userModule.NewUserModule(
		&conf,
		userModule.WithDB(dbManager.DB),
		userModule.WithHTTPApp(httpApp),
	)
	userModel.Run()
	httpApp.Listen(":8080")
}
