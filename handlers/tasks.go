package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kapit4n/state/models"
)

var tasks = []models.Task{
	{Title: "Title"},
	{Title: "Title 2"},
}

func GetTasks(c *gin.Context) {
	var tasksFromDB []models.Task

	models.DB.Find(&tasksFromDB)

	c.IndentedJSON(http.StatusOK, tasks)
}

func GetTaskById(c *gin.Context) {
	id := c.Param("id")
	var task models.Task
	models.DB.Where("id = ?", id).First(&task)

	c.IndentedJSON(http.StatusOK, task)
}
