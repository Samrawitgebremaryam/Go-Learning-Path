package database

import (
	"Task_Manager/domain"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoTaskDatabase struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewMongoTaskDatabase(client *mongo.Client) domain.TaskDatabase {
	return &MongoTaskDatabase{
		client:     client,
		collection: client.Database("taskManager").Collection("tasks"),
	}
}

func (mtd *MongoTaskDatabase) FindTasks(filter interface{}) ([]domain.Task, error) {
	var tasks []domain.Task
	cursor, err := mtd.collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	if err = cursor.All(context.TODO(), &tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}

func (mtd *MongoTaskDatabase) FindTaskByID(id primitive.ObjectID, filter interface{}) (domain.Task, error) {
	var task domain.Task
	err := mtd.collection.FindOne(context.TODO(), filter).Decode(&task)
	if err != nil {
		return task, err
	}
	return task, nil
}

func (mtd *MongoTaskDatabase) CreateTask(task domain.Task) error {
	_, err := mtd.collection.InsertOne(context.TODO(), task)
	return err
}

func (mtd *MongoTaskDatabase) UpdateTask(id primitive.ObjectID, update interface{}) error {
	result, err := mtd.collection.UpdateOne(context.TODO(), bson.M{"_id": id}, bson.M{"$set": update})
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return mongo.ErrNoDocuments
	}
	return nil
}

func (mtd *MongoTaskDatabase) DeleteTask(id primitive.ObjectID, filter interface{}) error {
	result, err := mtd.collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return mongo.ErrNoDocuments
	}
	return nil
}
