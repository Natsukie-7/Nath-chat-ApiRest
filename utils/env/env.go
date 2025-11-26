package Env

import (
	"github.com/Natsukie-7/Nath-chat-ApiRest/utils/env/EnvLoader"
)

func init() {
	loader := EnvLoader.Load()

	HOST = loader.String("API_HOST")
	PORT = loader.Int("API_PORT", 3000)
	DEBUG = loader.Bool("DEBUG", false)

	BCRYPT_KEYWORD = loader.String("BCRYPT_KEYWORD")
}

var (
	HOST           string
	PORT           int
	BCRYPT_KEYWORD string
	DEBUG          bool
)
