package UserController

import "github.com/gofiber/fiber/v3"

func GetUser(ctx fiber.Ctx) error {
	return ctx.JSON(fiber.Map{"data": "Nath"})
}
