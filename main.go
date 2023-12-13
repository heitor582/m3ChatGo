package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/heitor582/m3ChatGo/src/configuration"
	"github.com/heitor582/m3ChatGo/src/routes"
)

func init() {
	configuration.InitDatabase()
}

func main() {
	var PORT string = ":" + os.Getenv("PORT")
	app := fiber.New()
	routes.SetupAuthRoutes(app)
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(os.Getenv("JWT_SECRET_KEY")),
	}))
	routes.SetupChatRoomRoutes(app)
	routes.SetupChatRoutes(app)
	app.Listen(PORT)
}