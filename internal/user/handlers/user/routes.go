package userHandler

import "github.com/gofiber/fiber/v2"

func (u *UserHandler) SetupRoutes(httpApp *fiber.App) {
	u.httpApp = httpApp

	v1 := u.httpApp.Group("/v1")
	u.setupUserRoutes(v1)
}

func (u *UserHandler) setupUserRoutes(v1 fiber.Router) {
	userGroup := v1.Group("/users")
	userGroup.Get("", u.ListUser)
}
