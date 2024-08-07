package controller

import (
	"Task_Manager/data"
	"Task_Manager/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserController struct {
	service data.UserManager
}

func NewUserController(usermgr data.UserManager) *UserController {
	return &UserController{
		service: usermgr,
	}
}

func (controller *UserController) RegisterUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	status, err := controller.service.RegisterUser(user)
	if err != nil {
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

func (controller *UserController) LoginUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	status, err, token := controller.service.LoginUserDb(user)
	if err != nil {
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User logged in successfully", "token": token})
}

// DeleteTask handles DELETE requests to delete a user by its ID
func (controller *UserController) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	idint, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid task ID"})
		return
	}
	erro := controller.service.DeleteUser(idint)
	if erro != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": erro.Error()})
		return
	}
	c.IndentedJSON(http.StatusAccepted, gin.H{"message": "User deleted"})
}
