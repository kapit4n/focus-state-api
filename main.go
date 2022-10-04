package main

import (
	"github.com/gin-gonic/gin"

	"github.com/kapit4n/state/handlers"
)

func main() {
	router := gin.Default()
	router.GET("/tasks", handlers.GetTasks)
	router.GET("/tasks/:id", handlers.GetTaskById)

	router.Run("localhost:8080")
}
