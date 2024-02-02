package main

import (
	"heldesk-api/database"
	"heldesk-api/route"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	database.Connect()

	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New())

	route.SetupRoutes(app)

	// handle unvailable routes
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})

	app.Listen(":3000")
}
