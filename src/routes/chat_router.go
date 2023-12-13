package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/heitor582/m3ChatGo/src/controllers"
)

func SetupChatRoutes(app *fiber.App) {
	route := app.Group("/api/chat")
	route.Post("/", controllers.SendMessage)
}