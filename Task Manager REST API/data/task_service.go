package data

import (
	"Task_Manager/models"
	"errors"
	"time"
)

type TaskManager struct {
	tasks []models.Task
}

// Sample in-memory data for tasks
func NewTaskManager() *TaskManager {
	return &TaskManager{
		tasks: []models.Task{
			{ID: 1, Title: "Task 1", Description: "First task", Due_date: time.Now(), Status: "Pending"},
			{ID: 2, Title: "Task 2", Description: "Second task", Due_date: time.Now().AddDate(0, 0, 1), Status: "In Progress"},
			{ID: 3, Title: "Task 3", Description: "Third task", Due_date: time.Now().AddDate(0, 0, 2), Status: "Completed"},
		},
	}
}

// GetTasks returns a list of all tasks
func (taskmgr *TaskManager) GetTasks() []models.Task {
	return taskmgr.tasks
}

// GetTaskById returns the details of a specific task by its ID
func (taskmgr *TaskManager) GetTaskById(id int) (models.Task, error) {
	for _, task := range taskmgr.tasks {
		if id == task.ID {
			return task, nil
		}
	}
	return models.Task{}, errors.New("task not found")
}

// DeleteTask deletes a specific task by its ID
func (taskmgr *TaskManager) DeleteTask(id int) error {
	for i, task := range taskmgr.tasks {
		if id == task.ID {
			taskmgr.tasks = append(taskmgr.tasks[:i], taskmgr.tasks[i+1:]...)
			return nil
		}
	}
	return errors.New("task not found")
}

// PutTask updates the details of a specific task by its ID
func (taskmgr *TaskManager) PutTask(update models.Task, id int) error {
	for i, task := range taskmgr.tasks {
		if id == task.ID {
			taskmgr.tasks[i] = update
			return nil
		}
	}
	return errors.New("task not found")
}

// PostTask creates a new task
func (taskmgr *TaskManager) PostTask(newTask models.Task) error {
	for _, task := range taskmgr.tasks {
		if newTask.ID == task.ID {
			return errors.New("task ID already exists")
		}
	}
	taskmgr.tasks = append(taskmgr.tasks, newTask)
	return nil
}
