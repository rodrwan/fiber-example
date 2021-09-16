package users

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/rodrwan/fiber-example/internal/app"
	"github.com/rodrwan/fiber-example/internal/models/user"
)

func GetUsers(ctx app.Context) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		return c.SendString(fmt.Sprintf("And the API is UP! %s", ctx.DB))
	}
}

func CreateUser(ctx app.Context) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var u *user.User
		if err := c.BodyParser(u); err != nil {
			return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
		}

		return c.Status(201).JSON(fiber.Map{"status": "success", "message": "Created User", "data": u})
	}
}
