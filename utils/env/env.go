// utils/env/config.go
package Env

import (
	"github.com/Natsukie-7/Nath-chat-ApiRest/utils/env/EnvLoader"
)

func init() {
	loader := EnvLoader.Load()

	Port = loader.Int("PORT", 3000)
	BcryptKeyword = loader.String("BCRYPT_KEYWORD")
	Debug = loader.Bool("DEBUG", false)
	Teste = loader.String("Teste")
}

var (
	Port          int
	BcryptKeyword string
	Debug         bool
	Teste         string
)
