package router

import (
	"Task_Manager/controller"

	"github.com/gin-gonic/gin"
)

func Endpoints(r *gin.Engine, taskmgr *controller.TaskController) {
	r.GET("/tasks", taskmgr.GetTasks)          // Define a route to get a list of all tasks
	r.POST("/tasks", taskmgr.PostTask)         // Define a route to create a new task
	r.PUT("/tasks/:id", taskmgr.PutTask)       // Define a route to update a specific task by its ID
	r.GET("/tasks/:id", taskmgr.GetTaskById)   // Define a route to get the details of a specific task by its ID
	r.DELETE("/tasks/:id", taskmgr.DeleteTask) // Define a route to delete a specific task by its ID

}
