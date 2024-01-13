package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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
	app.Use(cors.New(cors.Config{
        AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
        AllowOrigins:     "*",
        AllowCredentials: true,
        AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
    }))
	routes.SetupAuthRoutes(app)
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(os.Getenv("JWT_SECRET_KEY")),
	}))
	routes.SetupChatRoomRoutes(app)
	routes.SetupChatRoutes(app)
	app.Listen(PORT)
}