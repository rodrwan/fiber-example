package users

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rodrwan/fiber-example/internal/app"
	usersHandler "github.com/rodrwan/fiber-example/internal/handlers/users"
)

func UsersEnpoints(ctx app.Context, api fiber.Router) {
	user := api.Group("/users")
	user.Get("/", usersHandler.GetUsers(ctx))
	user.Post("/", usersHandler.CreateUser(ctx))
}
