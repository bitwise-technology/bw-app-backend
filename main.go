package main

import (
	"bw-api/pkg/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	routes.HelloRoute(app)
	routes.NotFoundRoute(app)

	app.Listen(":3000")
}
