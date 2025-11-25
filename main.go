package main

import (
	"fmt"
	"log"

	"github.com/Natsukie-7/Nath-chat-ApiRest/routes"
	Api "github.com/Natsukie-7/Nath-chat-ApiRest/utils/api"
	Env "github.com/Natsukie-7/Nath-chat-ApiRest/utils/env"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
)

func main() {
	app := fiber.New()

	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{
			fiber.MethodGet,
			fiber.MethodPost,
			fiber.MethodPut,
			fiber.MethodDelete,
			fiber.MethodOptions,
		},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: false,
	}))

	app.Use(Api.SettupApiResponse())

	routes.SetupRoutes(app)

	welcomeBanner := fmt.Sprintf(`
	╭──────────────────────────────────────────╮
	│     Nath-chat API Inicializado!          │
	╰──────────────────────────────────────────╯
	Network: %s:%d
	`, Env.HOST, Env.PORT)
	fmt.Println(welcomeBanner)

	serverStartError := app.Listen(":3000", fiber.ListenConfig{
		DisableStartupMessage: true,
	})
	if serverStartError != nil {
		log.Fatalf("Erro ao iniciar servidor: %v", serverStartError)
	}
}
