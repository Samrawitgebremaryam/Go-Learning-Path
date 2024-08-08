package usecases

import (
	"Task_Manager/domain"
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskUsecase struct {
	taskRepo domain.TaskRepository
	timeout  time.Duration
}

func NewTaskUsecase(taskRepo domain.TaskRepository, timeout time.Duration) domain.TaskUsecase {
	return &TaskUsecase{
		taskRepo: taskRepo,
		timeout:  timeout,
	}
}

func (u *TaskUsecase) GetTasks(ctx context.Context, isAdmin string, userId primitive.ObjectID) ([]domain.Task, error) {
	_, cancel := context.WithTimeout(ctx, u.timeout)
	defer cancel()

	return u.taskRepo.GetTasks(isAdmin, userId)
}

func (u *TaskUsecase) GetTaskById(ctx context.Context, id primitive.ObjectID, isAdmin string, userId primitive.ObjectID) (domain.Task, error) {
	_, cancel := context.WithTimeout(ctx, u.timeout)
	defer cancel()

	return u.taskRepo.GetTaskById(id, isAdmin, userId)
}

func (u *TaskUsecase) DeleteTask(ctx context.Context, id primitive.ObjectID, isAdmin string, userId primitive.ObjectID) error {
	_, cancel := context.WithTimeout(ctx, u.timeout)
	defer cancel()

	return u.taskRepo.DeleteTask(id, isAdmin, userId)
}

func (u *TaskUsecase) PutTask(ctx context.Context, updatedTask domain.Task, id primitive.ObjectID, isAdmin string, userId primitive.ObjectID) error {
	_, cancel := context.WithTimeout(ctx, u.timeout)
	defer cancel()

	if updatedTask.Title == "" {
		return errors.New("title is required")
	}

	return u.taskRepo.PutTask(updatedTask, id, isAdmin, userId)
}

func (u *TaskUsecase) PostTask(ctx context.Context, newTask domain.Task) error {
	_, cancel := context.WithTimeout(ctx, u.timeout)
	defer cancel()

	if newTask.Title == "" {
		return errors.New("title is required")
	}

	return u.taskRepo.PostTask(newTask)
}
