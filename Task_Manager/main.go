package main

import (
	"Task_Manager/router" // Importing the router package for endpoint definitions
	"log"

	"github.com/gin-gonic/gin" // Importing the Gin framework for handling HTTP requests
)

func main() {
	// Initialize the Gin router with default middleware (logger and recovery middleware)
	r := gin.Default()
	// Register the application endpoints using the router package
	router.Endpoints(r)
	// Run the Gin HTTP server on the default port (8080)
	if err := r.Run(); err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}
