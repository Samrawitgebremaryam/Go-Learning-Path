package data

import (
	"Task_Manager/models"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskManager struct {
	client     *mongo.Client
	collection *mongo.Collection
}

// Sample in-memory data for tasks
func NewTaskManager(mclient *mongo.Client) *TaskManager {
	collection := mclient.Database("taskManager").Collection("tasks")
	return &TaskManager{
		client:     mclient,
		collection: collection,
	}
}

// GetTasks returns a list of all tasks
func (taskmgr *TaskManager) GetTasks() ([]models.Task, error) {
	var tasks []models.Task

	curser, err := taskmgr.collection.Find(context.TODO(), bson.D{})

	if err != nil {
		return nil, err
	}
	if err = curser.All(context.TODO(), &tasks); err != nil {
		return nil, err
	}
	return tasks, nil

}

// GetTaskById returns the details of a specific task by its ID
func (taskmgr *TaskManager) GetTaskById(id primitive.ObjectID) (models.Task, error) {
	var task models.Task

	filter := bson.D{{Key: "_id", Value: id}}

	err := taskmgr.collection.FindOne(context.TODO(), filter).Decode(&task)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return task, errors.New("task not found .invalid ID please input another ID")
		}
		return task, err
	}
	return task, nil
}

// DeleteTask deletes a specific task by its ID
func (taskmgr *TaskManager) DeleteTask(id primitive.ObjectID) error {

	filter := bson.D{{Key: "_id", Value: id}}

	result, err := taskmgr.collection.DeleteOne(context.TODO(), filter)

	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("task not found invalid ID please input another ID")
	}
	return nil
}

// PutTask updates the details of a specific task by its ID
func (taskmgr *TaskManager) PutTask(update models.Task, id primitive.ObjectID) error {
	filter := bson.D{{Key: "_id", Value: id}}
	updatebson := bson.D{{Key: "$set", Value: update}}
	result, err := taskmgr.collection.UpdateOne(context.TODO(), filter, updatebson)

	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return errors.New("task not found invalid ID please input another ID")

	}
	return nil
}

// PostTask creates a new task
func (taskmgr *TaskManager) PostTask(newTask models.Task) error {
	newTask.ID = primitive.NewObjectID()
	_, err := taskmgr.collection.InsertOne(context.TODO(), newTask)

	return err
}
