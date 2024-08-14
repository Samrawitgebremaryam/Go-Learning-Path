package database

import (
	"Task_Manager/domain"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoUserDatabase struct {
	collection *mongo.Collection
}

func NewMongoUserDatabase(client *mongo.Client) domain.UserDatabase {
	return &MongoUserDatabase{
		collection: client.Database("taskManager").Collection("users"),
	}
}

func (mud *MongoUserDatabase) FindUserByEmail(ctx context.Context, email string) (domain.User, error) {
	var user domain.User
	err := mud.collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (mud *MongoUserDatabase) CreateUser(ctx context.Context, user domain.User) error {
	_, err := mud.collection.InsertOne(ctx, user)
	return err
}

func (mud *MongoUserDatabase) DeleteUser(ctx context.Context, id primitive.ObjectID) error {
	result, err := mud.collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return mongo.ErrNoDocuments
	}
	return nil
}
