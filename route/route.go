package route

import (
	"heldesk-api/config"
	"heldesk-api/handler"
	"heldesk-api/middleware"

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

	/*
	* Auth routes
	 */

	//create a new JWT middleware
	jwt := middleware.NewAuthMiddleware(config.Secret)
	api.Post("/login", handler.Login)
	api.Get("/protected", jwt, handler.Protected)
}
