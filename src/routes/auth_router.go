package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/heitor582/m3ChatGo/src/controllers"
)

func SetupAuthRoutes(app *fiber.App) {
	route := app.Group("/api/auth")
	route.Post("/signin", controllers.Login)
	route.Post("/signup", controllers.RegisterUser)
}
