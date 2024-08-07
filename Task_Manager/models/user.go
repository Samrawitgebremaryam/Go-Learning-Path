package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID         primitive.ObjectID `bson:"_id"`
	First_name string             `json:"first_name" validate:"min=2 ,max = 100" `
	Last_name  string             `json:"last_name"  validate:"min=2 ,max = 100"`
	Password   string             `json:"password" validate:"required ,min=6"`
	Email      string             `json:"email" validate:"email, required"`
	Phone      string             `json:"phone"`
	User_type  string             `json:"user_type"`
}
