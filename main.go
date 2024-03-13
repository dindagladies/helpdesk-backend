package main

import (
	"user-api/cmd/handler"
	"user-api/cmd/storage"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", handler.Home)

	storage.InitDB()
	e.Logger.Fatal(e.Start(":3000"))
}
