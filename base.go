package main

import (
	"helpdesk/api/config"
	"helpdesk/api/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.DBInit()
	inDB := &controllers.InDB{DB: db}

	router := gin.Default()

	router.GET("/user/:id", inDB.Show)
	router.GET("/users", inDB.Index)
	router.POST("/user", inDB.Store)
	router.PUT("/user/:id", inDB.Update)
	router.DELETE("/user/:id", inDB.Delete)
	router.Run("localhost:3000")

	// run :
	// nodemon --watch './**/*.go' --signal SIGKILL --exec go run base.go
}
