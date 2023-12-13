package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/heitor582/m3ChatGo/src/controllers"
)

func SetupChatRoomRoutes(app *fiber.App) {
	route := app.Group("/api/chatrooms")
	route.Get("/", controllers.GetAllChatRooms)
	route.Get("/:id", controllers.GetChatRoom)
	route.Post("/", controllers.NewChatRoom)
	route.Delete("/:id", controllers.DeleteChatRoom)
}