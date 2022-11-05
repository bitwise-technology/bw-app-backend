package routes

import (
	"bw-api/internal/controllers"

	"github.com/gofiber/fiber/v2"
)

func HelloRoute(app *fiber.App) {
	app.Get("/", controllers.HelloController)
}
