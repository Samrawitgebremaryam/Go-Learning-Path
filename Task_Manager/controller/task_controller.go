package controller

import (
	"Task_Manager/data"
	"Task_Manager/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetTasks handles GET requests to retrieve all tasks
func GetTasks(c *gin.Context) {
	tasks := data.GetTasks()
	c.IndentedJSON(http.StatusOK, tasks)
}

// GetTaskById handles GET requests to retrieve a task by its ID
func GetTaskById(c *gin.Context) {
	id := c.Param("id")
	idint, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid task ID"})
		return
	}

	task, err := data.GetTaskById(idint)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, task)
}

// DeleteTask handles DELETE requests to delete a task by its ID
func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	idint, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid task ID"})
		return
	}

	err = data.DeleteTask(idint)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusAccepted, gin.H{"message": "Task deleted"})
}

// PutTask handles PUT requests to update a task by its ID
func PutTask(c *gin.Context) {
	id := c.Param("id")
	var updatableTask models.Task
	if err := c.BindJSON(&updatableTask); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON"})
		return
	}

	idint, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid task ID"})
		return
	}

	err = data.PutTask(updatableTask, idint)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusAccepted, gin.H{"message": "Task updated"})
}

// PostTask handles POST requests to create a new task
func PostTask(c *gin.Context) {
	var newTask models.Task
	if err := c.BindJSON(&newTask); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON"})
		return
	}

	err := data.PostTask(newTask)
	if err != nil {
		c.IndentedJSON(http.StatusConflict, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusCreated, gin.H{"message": "Task added"})
}
