package UserController

import (
	Api "github.com/Natsukie-7/Nath-chat-ApiRest/utils"
	"github.com/gofiber/fiber/v3"
)

func GetUser(ctx fiber.Ctx) error {
	return Api.Resp{
		Status:  201,
		Message: "Teste6",
		Data: Api.DataStruct{
			"User": "nath",
		},
	}
}
