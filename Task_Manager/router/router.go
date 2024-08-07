package router

import (
	"Task_Manager/controller"
	"Task_Manager/middleware"

	"github.com/gin-gonic/gin"
)

func Endpoints(r *gin.Engine, taskmgr *controller.TaskController, usermgr *controller.UserController) {
	auth := r.Group("/")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.GET("/tasks", taskmgr.GetTasks)        // Define a route to get a list of all tasks
		auth.PUT("/tasks/:id", taskmgr.PutTask)     // Define a route to update a specific task by its ID
		auth.GET("/tasks/:id", taskmgr.GetTaskById) // Define a route to get the details of a specific task by its ID

		admin := auth.Group("/")
		admin.Use(middleware.AdminMiddleware())
		{
			admin.POST("/tasks", taskmgr.PostTask)         // Define a route to create a new task
			admin.DELETE("/tasks/:id", taskmgr.DeleteTask) // Define a route to delete a specific task by its ID
		}

	}
	r.POST("/register", usermgr.RegisterUser)
	r.POST("/login", usermgr.LoginUser)
}
