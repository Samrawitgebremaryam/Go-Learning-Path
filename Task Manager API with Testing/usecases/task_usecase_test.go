package usecases

import (
	"Task_Manager/domain"
	"Task_Manager/mocks/repomocks"
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskUsecaseTestSuite struct {
	suite.Suite
	TaskRepoMock *repomocks.TaskRepository
	TaskUsecase  domain.TaskUsecase
}

func (s *TaskUsecaseTestSuite) SetupTest() {
	s.TaskRepoMock = new(repomocks.TaskRepository)
	s.TaskUsecase = NewTaskUsecase(s.TaskRepoMock, 5*time.Minute)
}

func (s *TaskUsecaseTestSuite) TearDownTest() {
	s.TaskRepoMock.AssertExpectations(s.T())
}

func (s *TaskUsecaseTestSuite) TestGetTasks() {
	userID := primitive.NewObjectID()
	isAdmin := "Admin"
	expectedTasks := []domain.Task{
		{ID: primitive.NewObjectID(), Title: "Task 1", Description: "Prepare the slides.",
			Due_date: time.Date(2024, time.August, 4, 12, 15, 48, 58, time.UTC),
			Status:   "In Progress"},
		{ID: primitive.NewObjectID(), Title: "Task 2", Description: "Prepare the slides.",
			Due_date: time.Date(2024, time.August, 4, 12, 15, 48, 58, time.UTC),
			Status:   "Done"},
	}

	s.TaskRepoMock.On("GetTasks", isAdmin, userID).Return(expectedTasks, nil).Once()

	tasks, err := s.TaskUsecase.GetTasks(context.Background(), isAdmin, userID)
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), expectedTasks, tasks)
}

func (s *TaskUsecaseTestSuite) TestAddTask() {
	newTask := domain.Task{
		ID:          primitive.NewObjectID(),
		UserId:      primitive.NewObjectID(),
		Title:       "Prepare hi hi ho",
		Description: "Prepare the slides.",
		Due_date:    time.Date(2024, time.August, 4, 12, 15, 48, 58, time.UTC),
		Status:      "In Progress",
	}

	s.TaskRepoMock.On("PostTask", mock.AnythingOfType("domain.Task")).Run(func(args mock.Arguments) {
		actualTask := args.Get(0).(domain.Task)
		s.T().Log("PostTask called with:", actualTask)
	}).Return(nil).Once()

	err := s.TaskUsecase.PostTask(context.Background(), newTask)
	assert.Nil(s.T(), err)

	s.TaskRepoMock.AssertExpectations(s.T())
}

func (s *TaskUsecaseTestSuite) TestDeleteTask() {
	taskID := primitive.NewObjectID()
	userID := primitive.NewObjectID()
	isAdmin := "true"

	s.TaskRepoMock.On("DeleteTask", taskID, isAdmin, userID).Return(nil).Once()

	err := s.TaskUsecase.DeleteTask(context.Background(), taskID, isAdmin, userID)
	assert.Nil(s.T(), err)

	s.TaskRepoMock.AssertExpectations(s.T())
}

func (s *TaskUsecaseTestSuite) TestGetTaskById() {
	taskID := primitive.NewObjectID()
	userID := primitive.NewObjectID()
	isAdmin := "true"
	expectedTask := domain.Task{
		ID:          taskID,
		UserId:      userID,
		Title:       "Sample Task",
		Description: "This is a sample task for testing.",
		Due_date:    time.Now(),
		Status:      "In Progress",
	}

	s.Run("task found", func() {
		s.TaskRepoMock.On("GetTaskById", taskID, isAdmin, userID).Return(expectedTask, nil).Once()
		task, err := s.TaskUsecase.GetTaskById(context.Background(), taskID, isAdmin, userID)
		assert.Nil(s.T(), err)
		assert.Equal(s.T(), expectedTask, task)
		s.TaskRepoMock.AssertExpectations(s.T())
	})
}

func (s *TaskUsecaseTestSuite) TestPutTask() {
	taskID := primitive.NewObjectID()
	userID := primitive.NewObjectID()
	isAdmin := "true"
	updatedTask := domain.Task{
		ID:          taskID,
		UserId:      userID,
		Title:       "Updated Task",
		Description: "Updated description for existing task.",
		Due_date:    time.Now(),
		Status:      "Completed",
	}

	s.Run("update success", func() {
		s.TaskRepoMock.On("PutTask", updatedTask, taskID, isAdmin, userID).Return(nil).Once()
		err := s.TaskUsecase.PutTask(context.Background(), updatedTask, taskID, isAdmin, userID)
		assert.Nil(s.T(), err)
		s.TaskRepoMock.AssertExpectations(s.T())
	})

	s.Run("update failure", func() {
		s.TaskRepoMock.On("PutTask", updatedTask, taskID, isAdmin, userID).Return(errors.New("update failed")).Once()
		err := s.TaskUsecase.PutTask(context.Background(), updatedTask, taskID, isAdmin, userID)
		assert.Error(s.T(), err)
		s.TaskRepoMock.AssertExpectations(s.T())
	})

	s.Run("task not found", func() {
		s.TaskRepoMock.On("GetTaskById", taskID, isAdmin, userID).Return(domain.Task{}, errors.New("task not found")).Once()
		task, err := s.TaskUsecase.GetTaskById(context.Background(), taskID, isAdmin, userID)
		assert.Error(s.T(), err)
		assert.Equal(s.T(), domain.Task{}, task)
		s.TaskRepoMock.AssertExpectations(s.T())
	})
}

func TestTaskUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(TaskUsecaseTestSuite))
}
