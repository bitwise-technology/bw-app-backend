package fiber

import (
	"bw-api/pkg/routes"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func Run(port int) {
	app := fiber.New()

	routes.HelloRoute(app)
	routes.NotFoundRoute(app)

	// Listen to port 3000.
	log.Fatal(app.Listen(fmt.Sprintf(":%d", port)))
}
