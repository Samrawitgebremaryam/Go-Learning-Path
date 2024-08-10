package router

import (
	"Task_Manager/Delivery/controller"
	infrastructure "Task_Manager/infrastructure"

	"github.com/gin-gonic/gin"
)

func Endpoints(r *gin.Engine, taskmgr *controller.TaskController, usermgr *controller.UserController) {
	auth := r.Group("/")
	auth.Use(infrastructure.AuthMiddleware())
	{
		auth.GET("/tasks", taskmgr.GetTasks)
		auth.PUT("/tasks/:id", taskmgr.PutTask)
		auth.GET("/tasks/:id", taskmgr.GetTaskById)
		auth.POST("/tasks", taskmgr.PostTask)
		auth.DELETE("/tasks/:id", taskmgr.DeleteTask)

		admin := auth.Group("/")
		admin.Use(infrastructure.AdminMiddleware())

		{
			admin.DELETE("/users/:id", usermgr.DeleteUser) // Define a route to delete a specific user by its ID
		}

	}
	r.POST("/register", usermgr.RegisterUser)
	r.POST("/login", usermgr.LoginUser)
}
