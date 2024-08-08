package controller

import (
	"Task_Manager/domain"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskController struct {
	taskUsecase domain.TaskUsecase
}

func NewTaskController(taskUsecase domain.TaskUsecase) *TaskController {
	return &TaskController{
		taskUsecase: taskUsecase,
	}
}

// GetTasks handles GET requests to retrieve all tasks
func (controller *TaskController) GetTasks(c *gin.Context) {
	isAdmin := c.GetString("usertype")
	userID := c.GetString("userid")
	userId, _ := primitive.ObjectIDFromHex(userID)

	tasks, err := controller.taskUsecase.GetTasks(c, isAdmin, userId)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "couldn't fetch the data"})
		return
	}
	c.IndentedJSON(http.StatusOK, tasks)
}

// GetTaskById handles GET requests to retrieve a task by its ID
func (controller *TaskController) GetTaskById(c *gin.Context) {
	isAdmin := c.GetString("usertype")
	userID := c.GetString("userid")
	userId, _ := primitive.ObjectIDFromHex(userID)

	id := c.Param("id")
	idint, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid task ID"})
		return
	}

	task, err := controller.taskUsecase.GetTaskById(c, idint, isAdmin, userId)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, task)
}

// DeleteTask handles DELETE requests to delete a task by its ID
func (controller *TaskController) DeleteTask(c *gin.Context) {
	isAdmin := c.GetString("usertype")
	userID := c.GetString("userid")
	userId, _ := primitive.ObjectIDFromHex(userID)

	id := c.Param("id")
	idint, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid task ID"})
		return
	}

	err = controller.taskUsecase.DeleteTask(c, idint, isAdmin, userId)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusAccepted, gin.H{"message": "Task deleted"})
}

// PutTask handles PUT requests to update a task by its ID
func (controller *TaskController) PutTask(c *gin.Context) {
	isAdmin := c.GetString("usertype")
	userID := c.GetString("userid")
	userId, _ := primitive.ObjectIDFromHex(userID)
	id := c.Param("id")
	var updatableTask domain.Task
	if err := c.BindJSON(&updatableTask); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON"})
		return
	}

	idint, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid task ID"})
		return
	}

	err = controller.taskUsecase.PutTask(c, updatableTask, idint, isAdmin, userId)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusAccepted, gin.H{"message": "Task updated"})
}

// PostTask handles POST requests to create a new task
func (controller *TaskController) PostTask(c *gin.Context) {
	var newTask domain.Task
	userId := c.GetString("userid")
	newTask.ID = primitive.NewObjectID()
	userID, err := primitive.ObjectIDFromHex(userId)
	newTask.UserId = userID

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON"})
		return
	}

	if err := c.BindJSON(&newTask); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON"})
		return
	}
	erro := controller.taskUsecase.PostTask(c, newTask)
	if erro != nil {
		c.IndentedJSON(http.StatusConflict, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusCreated, gin.H{"message": "Task added"})
}
