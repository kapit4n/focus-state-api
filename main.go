package main

import (
	"github.com/gin-gonic/gin"

	"github.com/kapit4n/state/handlers"
	"github.com/kapit4n/state/models"
)

func main() {

	models.ConnectDatabase()

	// create new task
	models.DB.Create(&models.Task{Title: "Task 001"})

	router := gin.Default()
	router.GET("/tasks", handlers.GetTasks)
	router.GET("/tasks/:id", handlers.GetTaskById)

	router.Run("localhost:8080")
}
