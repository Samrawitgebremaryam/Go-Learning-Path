package models

import (
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
