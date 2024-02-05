package route

import (
	"heldesk-api/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// grouping
	api := app.Group("/api")
	v1 := api.Group("user")

	// routes
	v1.Get("/", handler.GetAllUsers)
	v1.Get("/:id", handler.GetUser)
	v1.Post("/", handler.CreateUser)
	v1.Put("/:id", handler.UpdateUser)
	v1.Delete("/:id", handler.DeleteUserByID)
}
