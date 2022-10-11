package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kapit4n/state/models"
)

var tasks = []models.Task{
	{ID: "1", Title: "Title"},
	{ID: "2", Title: "Title 2"},
}

func GetTasks(c *gin.Context) {
	var tasksFromDB []models.Task

	models.DB.Find(&tasksFromDB)

	c.IndentedJSON(http.StatusOK, tasks)
}

func GetTaskById(c *gin.Context) {
	task, err := filterTask(c.Param("id"))

	if err != nil {
		c.IndentedJSON(http.StatusOK, err)
		return
	}
	c.IndentedJSON(http.StatusOK, task)

}

func filterTask(id string) (*models.Task, error) {
	for _, t := range tasks {
		if t.ID == id {
			return &t, nil
		}
	}

	return nil, errors.New("Not found")
}
