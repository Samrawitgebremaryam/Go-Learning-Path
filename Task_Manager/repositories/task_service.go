package repositories

import (
	"Task_Manager/domain"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskRepository struct {
	client     *mongo.Client
	collection *mongo.Collection
}

// Sample in-memory data for tasks
func NewTaskRepository(mclient *mongo.Client) domain.TaskRepository {
	collection := mclient.Database("taskManager").Collection("tasks")
	return &TaskRepository{
		client:     mclient,
		collection: collection,
	}
}

// GetTasks returns a list of all tasks
func (taskmgr *TaskRepository) GetTasks(isAdmin string, userId primitive.ObjectID) ([]domain.Task, error) {
	var tasks []domain.Task

	filter := bson.D{}

	if isAdmin != "Admin" {
		filter = bson.D{{Key: "_userid", Value: userId}}
	}

	curser, err := taskmgr.collection.Find(context.TODO(), filter)

	if err != nil {
		return nil, err
	}
	if err = curser.All(context.TODO(), &tasks); err != nil {
		return nil, err
	}
	return tasks, nil

}

// GetTaskById returns the details of a specific task by its ID
func (taskmgr *TaskRepository) GetTaskById(id primitive.ObjectID, isAdmin string, userId primitive.ObjectID) (domain.Task, error) {
	var task domain.Task
	filter := bson.D{{Key: "_id", Value: id}}

	if isAdmin != "Admin" {
		filter = bson.D{
			{Key: "_userid", Value: userId},
			{Key: "_id", Value: id},
		}
	}

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
func (taskmgr *TaskRepository) DeleteTask(id primitive.ObjectID, isAdmin string, userId primitive.ObjectID) error {
	filter := bson.D{{Key: "_id", Value: id}}

	if isAdmin != "Admin" {
		filter = bson.D{
			{Key: "_userid", Value: userId},
			{Key: "_id", Value: id},
		}
	}

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
func (taskmgr *TaskRepository) PutTask(updatedTask domain.Task, id primitive.ObjectID, isAdmin string, userId primitive.ObjectID) error {

	filter := bson.D{{Key: "_id", Value: id}}

	if isAdmin != "Admin" {
		filter = bson.D{
			{Key: "_userid", Value: userId},
			{Key: "_id", Value: id},
		}
	}
	update := bson.M{
		"title":       updatedTask.Title,
		"description": updatedTask.Description,
		"due_date":    updatedTask.Due_date,
		"status":      updatedTask.Status,
	}

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
func (taskmgr *TaskRepository) PostTask(newTask domain.Task) error {

	_, erro := taskmgr.collection.InsertOne(context.TODO(), newTask)
	return erro
}
