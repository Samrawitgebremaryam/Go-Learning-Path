package main

import (
	"Task_Manager/controller"
	"Task_Manager/data"
	"Task_Manager/router"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	tasks := data.NewTaskManager()
	taskmgr := controller.NewTaskController(*tasks)

	r := gin.Default()
	router.Endpoints(r, taskmgr)

	if err := r.Run(); err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}
