package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/rodrwan/fiber-example/internal/app"
	usersRouter "github.com/rodrwan/fiber-example/internal/routes/users"
)

func main() {
	server := fiber.New()

	ctx := app.Context{
		DB: "postgres",
	}

	api := server.Group("/api", logger.New())
	usersRouter.UsersEnpoints(ctx, api)

	server.Listen(":8080")
}
