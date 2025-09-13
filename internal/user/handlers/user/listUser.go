package userHandler

import (
	"github.com/fbriansyah/go-modular/internal/model"
	userModel "github.com/fbriansyah/go-modular/internal/model/user"
	"github.com/gofiber/fiber/v2"
)

func (u *UserHandler) ListUser(c *fiber.Ctx) error {
	ctx := c.Context()

	query := &userModel.ListUserQuery{
		FirstName: c.Query("first_name"),
		LastName:  c.Query("last_name"),
		Email:     c.Query("email"),
		GeneralListQuery: model.GeneralListQuery{
			Limit:  c.QueryInt("limit", 10),
			Offset: c.QueryInt("offset", 0),
		},
	}

	users, err := u.userService.ListUser(ctx, query)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(users)
}
