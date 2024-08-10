package main

import (
	"Task_Manager/controller"
	"Task_Manager/data"
	"Task_Manager/router"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Retrieve the MongoDB URI from the environment variable
	MONGO_URI := os.Getenv("MONGO_URI")

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().ApplyURI(MONGO_URI).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	tasks := data.NewTaskManager(client)
	taskmgr := controller.NewTaskController(*tasks)
	users := data.NewUserManager(client)
	usermgr := controller.NewUserController(*users)

	r := gin.Default()
	router.Endpoints(r, taskmgr, usermgr)

	if err := r.Run(); err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}
