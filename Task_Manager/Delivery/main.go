package main

import (
	"Task_Manager/Delivery/controller"
	"Task_Manager/Delivery/router"
	"Task_Manager/repositories"
	"Task_Manager/usecases"
	"context"
	"fmt"
	"log"
	"os"
	"time"

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
	tasksrepo := repositories.NewTaskRepository(client)
	usersrepo := repositories.NewUserRepository(client)

	taskuse := usecases.NewTaskUsecase(tasksrepo, 60*time.Second)
	useruse := usecases.NewUserUsecase(usersrepo, 60*time.Second)

	taskcont := controller.NewTaskController(taskuse)
	usercont := controller.NewUserController(useruse)

	r := gin.Default()
	router.Endpoints(r, taskcont, usercont)

	if err := r.Run(); err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}
