package controllers

import "github.com/gofiber/fiber/v2"

func HelloController(context *fiber.Ctx) error {
	return context.JSON(fiber.Map{
		"error":   false,
		"message": "Hello, World!",
	})
}
