package routes

import "github.com/gofiber/fiber/v2"

func NotFoundRoute(app *fiber.App) {
	app.Use(
		func(context *fiber.Ctx) error {
			return context.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error":   true,
				"message": "Sorry, endpoint is not found",
			})
		},
	)
}
