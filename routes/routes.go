package routes

import (
	UserController "github.com/Natsukie-7/Nath-chat-ApiRest/controllers"
	"github.com/gofiber/fiber/v3"
)

func SetupRoutes(fiberApp *fiber.App) {
	app := fiberApp.Group("/api")

	app.Get("/", UserController.GetUser)
}
