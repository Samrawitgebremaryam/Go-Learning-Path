package domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	UserId      primitive.ObjectID `json:"userid" bson:"_userid"`
	Title       string             `json:"title"`
	Description string             `json:"description"`
	Due_date    time.Time          `json:"due_date"`
	Status      string             `json:"status"`
}

type User struct {
	ID         primitive.ObjectID `bson:"_id"`
	First_name string             `json:"first_name" validate:"min=2 ,max = 100" `
	Last_name  string             `json:"last_name"  validate:"min=2 ,max = 100"`
	Password   string             `json:"password" validate:"required ,min=6"`
	Email      string             `json:"email" validate:"email, required"`
	Phone      string             `json:"phone"`
	User_type  string             `json:"user_type"`
}
type TaskUsecase interface {
	GetTasks(ctx context.Context, isAdmin string, userId primitive.ObjectID) ([]Task, error)
	GetTaskById(ctx context.Context, id primitive.ObjectID, isAdmin string, userId primitive.ObjectID) (Task, error)
	DeleteTask(ctx context.Context, id primitive.ObjectID, isAdmin string, userId primitive.ObjectID) error
	PutTask(ctx context.Context, updatedTask Task, id primitive.ObjectID, isAdmin string, userId primitive.ObjectID) error
	PostTask(ctx context.Context, newTask Task) error
}

type TaskRepository interface {
	GetTasks(isAdmin string, userId primitive.ObjectID) ([]Task, error)
	GetTaskById(id primitive.ObjectID, isAdmin string, userId primitive.ObjectID) (Task, error)
	DeleteTask(id primitive.ObjectID, isAdmin string, userId primitive.ObjectID) error
	PutTask(updatedTask Task, id primitive.ObjectID, isAdmin string, userId primitive.ObjectID) error
	PostTask(newTask Task) error
}

type UserRepository interface {
	RegisterUser(user User) (int, error)
	LoginUser(user User) (int, error, string)
	DeleteUser(id primitive.ObjectID) error
}

type UserUsecase interface {
	RegisterUser(ctx context.Context, user User) (int, error)
	LoginUser(ctx context.Context, user User) (int, error, string)
	DeleteUser(ctx context.Context, id primitive.ObjectID) error
}
