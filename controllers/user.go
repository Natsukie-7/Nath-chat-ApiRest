package UserController

import (
	Api "github.com/Natsukie-7/Nath-chat-ApiRest/utils/api"
	"github.com/Natsukie-7/Nath-chat-ApiRest/utils/bycrypt"
	Env "github.com/Natsukie-7/Nath-chat-ApiRest/utils/env"
	"github.com/gofiber/fiber/v3"
)

func GetUser(ctx fiber.Ctx) error {
	pass, _ := bycrypt.HashPassword(Env.BCRYPT_KEYWORD)

	return Api.Resp{
		Status:  201,
		Message: "Sucesso",
		Data: Api.DataStruct{
			"User": pass,
		},
	}
}
