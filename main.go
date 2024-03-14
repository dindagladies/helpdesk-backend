package main

import (
	"user-api/cmd/handlers"
	"user-api/cmd/storage"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", handlers.Home)
	storage.InitDB()

	e.POST("/users", handlers.CreateUser)

	e.Logger.Fatal(e.Start(":3000"))
}
