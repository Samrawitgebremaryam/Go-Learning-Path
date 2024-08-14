package repositories

import (
	"Task_Manager/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskRepository struct {
	taskDB domain.TaskDatabase
}

func NewTaskRepository(taskDB domain.TaskDatabase) domain.TaskRepository {
	return &TaskRepository{
		taskDB: taskDB,
	}
}

func (tr *TaskRepository) GetTasks(isAdmin string, userId primitive.ObjectID) ([]domain.Task, error) {
	filter := bson.M{}
	if isAdmin != "Admin" {
		filter = bson.M{"_userid": userId}
	}
	return tr.taskDB.FindTasks(filter)
}

func (tr *TaskRepository) GetTaskById(id primitive.ObjectID, isAdmin string, userId primitive.ObjectID) (domain.Task, error) {
	filter := bson.M{"_id": id}
	if isAdmin != "Admin" {
		filter["_userid"] = userId
	}
	return tr.taskDB.FindTaskByID(id, filter)
}

func (tr *TaskRepository) DeleteTask(id primitive.ObjectID, isAdmin string, userId primitive.ObjectID) error {
	filter := bson.M{"_id": id}
	if isAdmin != "Admin" {
		filter["_userid"] = userId
	}
	return tr.taskDB.DeleteTask(id, filter)
}

func (tr *TaskRepository) PutTask(updatedTask domain.Task, id primitive.ObjectID, isAdmin string, userId primitive.ObjectID) error {
	update := bson.M{
		"title":       updatedTask.Title,
		"description": updatedTask.Description,
		"due_date":    updatedTask.Due_date,
		"status":      updatedTask.Status,
	}
	return tr.taskDB.UpdateTask(id, update)
}

func (tr *TaskRepository) PostTask(newTask domain.Task) error {
	return tr.taskDB.CreateTask(newTask)
}
