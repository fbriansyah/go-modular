package http

import (
	"os"
	"os/signal"

	"github.com/fbriansyah/go-modular/config"
	userModule "github.com/fbriansyah/go-modular/internal/user"
	"github.com/fbriansyah/go-modular/pkg/database"
	"github.com/gofiber/fiber/v2"
)

func Main() {
	conf, secret, err := config.LoadConfig("./config.yaml", "./secret.yaml")
	if err != nil {
		panic(err)
	}
	dbManager, err := database.NewManager(&conf.Database, &secret.Database, "./migrations")
	if err != nil {
		panic(err)
	}
	httpApp := fiber.New()
	userModel := userModule.NewUserModule(
		conf,
		userModule.WithDB(dbManager.DB),
		userModule.WithHTTPApp(httpApp),
	)
	userModel.Run()
	httpApp.Listen(":8080")

	// setup graceful shutdown, listen for interrupt signal
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	<-signalChan
	httpApp.Shutdown()
}
