package UserController

import (
	Api "github.com/Natsukie-7/Nath-chat-ApiRest/utils/api"
	"github.com/Natsukie-7/Nath-chat-ApiRest/utils/bycrypt"
	"github.com/gofiber/fiber/v3"
)

func GetUser(ctx fiber.Ctx) error {
	pass, _ := bycrypt.HashPassword("Nath")

	return Api.Resp{
		Status:  201,
		Message: "Sucesso",
		Data: Api.DataStruct{
			"User": pass,
		},
	}
}
