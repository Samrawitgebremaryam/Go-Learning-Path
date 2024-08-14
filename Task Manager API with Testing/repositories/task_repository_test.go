package repositories

import (
	"Task_Manager/domain"
	"Task_Manager/mocks/database_mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TaskRepositoryTestsstruct for holding test setup
type TaskRepositoryTestsSuite struct {
	suite.Suite
	mockDB     *database_mocks.TaskDatabase
	repository domain.TaskRepository
}

func (s *TaskRepositoryTestsSuite) SetupTest() {
	s.mockDB = new(database_mocks.TaskDatabase) // Initialize the mock
	s.repository = NewTaskRepository(s.mockDB)
}

func (s *TaskRepositoryTestsSuite) TestGetTasks() {
	userID := primitive.NewObjectID()
	tasks := []domain.Task{
		{ID: primitive.NewObjectID(), Title: "Task 1", Description: "Description 1"},
		{ID: primitive.NewObjectID(), Title: "Task 2", Description: "Description 2"},
	}

	filter := bson.M{"_userid": userID}
	s.mockDB.On("FindTasks", filter).Return(tasks, nil).Once() // Setup expectation

	result, err := s.repository.GetTasks("notAdmin", userID) // Call the method

	assert.NoError(s.T(), err)
	assert.Equal(s.T(), tasks, result) // Assert correct data received
	s.mockDB.AssertExpectations(s.T()) // Assert that all expectations were met
}

func (s *TaskRepositoryTestsSuite) TestGetTaskById() {
	userID := primitive.NewObjectID()
	taskID := primitive.NewObjectID()
	expectedTask := domain.Task{ID: taskID, Title: "Specific Task"}

	filter := bson.M{"_id": taskID, "_userid": userID}
	s.mockDB.On("FindTaskByID", taskID, filter).Return(expectedTask, nil).Once()

	task, err := s.repository.GetTaskById(taskID, "notAdmin", userID)

	assert.NoError(s.T(), err)
	assert.Equal(s.T(), expectedTask, task)
	s.mockDB.AssertExpectations(s.T())
}

func (s *TaskRepositoryTestsSuite) TestPutTask() {
	userID := primitive.NewObjectID()
	taskID := primitive.NewObjectID()
	updatedTask := domain.Task{
		ID:          taskID,
		Title:       "Updated Task",
		Description: "Updated Description",
		Due_date:    time.Now(),
		Status:      "Completed",
	}

	update := bson.M{
		"title":       updatedTask.Title,
		"description": updatedTask.Description,
		"due_date":    updatedTask.Due_date,
		"status":      updatedTask.Status,
	}
	s.mockDB.On("UpdateTask", taskID, update).Return(nil).Once()

	err := s.repository.PutTask(updatedTask, taskID, "Admin", userID)

	assert.NoError(s.T(), err)
	s.mockDB.AssertExpectations(s.T())
}

func (s *TaskRepositoryTestsSuite) TestPostTask() {
	newTask := domain.Task{
		Title:       "New Task",
		Description: "Task description",
		Due_date:    time.Now(),
		Status:      "Pending",
	}

	s.mockDB.On("CreateTask", newTask).Return(nil).Once()

	err := s.repository.PostTask(newTask)

	assert.NoError(s.T(), err)
	s.mockDB.AssertExpectations(s.T())
}

func (s *TaskRepositoryTestsSuite) TestDeleteTask() {
	userID := primitive.NewObjectID()
	taskID := primitive.NewObjectID()

	filter := bson.M{"_id": taskID, "_userid": userID}
	s.mockDB.On("DeleteTask", taskID, filter).Return(nil).Once()

	err := s.repository.DeleteTask(taskID, "notAdmin", userID)

	assert.NoError(s.T(), err)
	s.mockDB.AssertExpectations(s.T())
}

// Run the test s
func TestTaskRepositoryTestsSuite(t *testing.T) {
	suite.Run(t, new(TaskRepositoryTestsSuite))
}
