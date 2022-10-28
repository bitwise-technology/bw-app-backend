package routes

import (
	"bw-api/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func HelloRoute(app *fiber.App) {
	app.Get("/", controllers.HelloController)
}
