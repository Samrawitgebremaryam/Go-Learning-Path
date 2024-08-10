package controller

import (
	"Task_Manager/data"
	"Task_Manager/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskController struct {
	service data.TaskManager
}

func NewTaskController(taskmgr data.TaskManager) *TaskController {
	return &TaskController{
		service: taskmgr,
	}
}

// GetTasks handles GET requests to retrieve all tasks
func (controller *TaskController) GetTasks(c *gin.Context) {
	tasks, err := controller.service.GetTasks()
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "couldn't fetch the data"})
		return
	}
	c.IndentedJSON(http.StatusOK, tasks)
}

// GetTaskById handles GET requests to retrieve a task by its ID
func (controller *TaskController) GetTaskById(c *gin.Context) {
	id := c.Param("id")
	idint, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid task ID"})
		return
	}

	task, err := controller.service.GetTaskById(idint)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, task)
}

// DeleteTask handles DELETE requests to delete a task by its ID
func (controller *TaskController) DeleteTask(c *gin.Context) {
	id := c.Param("id")
	idint, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid task ID"})
		return
	}

	err = controller.service.DeleteTask(idint)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusAccepted, gin.H{"message": "Task deleted"})
}

// PutTask handles PUT requests to update a task by its ID
func (controller *TaskController) PutTask(c *gin.Context) {
	id := c.Param("id")
	var updatableTask models.Task
	if err := c.BindJSON(&updatableTask); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON"})
		return
	}

	idint, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid task ID"})
		return
	}

	err = controller.service.PutTask(updatableTask, idint)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusAccepted, gin.H{"message": "Task updated"})
}

// PostTask handles POST requests to create a new task
func (controller *TaskController) PostTask(c *gin.Context) {
	var newTask models.Task
	if err := c.BindJSON(&newTask); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON"})
		return
	}

	err := controller.service.PostTask(newTask)
	if err != nil {
		c.IndentedJSON(http.StatusConflict, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusCreated, gin.H{"message": "Task added"})
}
