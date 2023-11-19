package routes

import (
	"github.com/goravel/framework/facades"

	"goravel/app/http/controllers"
)

func Api() {
	userController := controllers.NewUserController()
	facades.Route().Get("/users", userController.Index)
	facades.Route().Get("/users/{id}", userController.Show)
	facades.Route().Post("/users", userController.Store)
	facades.Route().Put("/users/{id}", userController.Update)
	facades.Route().Delete("/users/{id}", userController.Destroy)
}
