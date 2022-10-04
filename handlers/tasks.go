package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kapit4n/state/models"
)

var tasks = []models.Task{
	{ID: "1", Title: "Title"},
}

func GetTasks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, tasks)
}
