package controller

import (
	"Task_Manager/domain"
	"Task_Manager/mocks/usecase_mocks"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskControllerTestSuite struct {
	suite.Suite
	router      *gin.Engine
	taskUsecase *usecase_mocks.TaskUsecase
}

func (suite *TaskControllerTestSuite) SetupTest() {
	gin.SetMode(gin.TestMode)
	suite.router = gin.Default()
	suite.taskUsecase = new(usecase_mocks.TaskUsecase)
	taskController := NewTaskController(suite.taskUsecase)
	suite.router.GET("/tasks", taskController.GetTasks)
	suite.router.GET("/tasks/:id", taskController.GetTaskById)
	suite.router.POST("/tasks", taskController.PostTask)
	suite.router.PUT("/tasks/:id", taskController.PutTask)
	suite.router.DELETE("/tasks/:id", taskController.DeleteTask)
}

func (suite *TaskControllerTestSuite) TestGetTasks() {
	tasks := []domain.Task{{ID: primitive.NewObjectID(), Title: "Task 1"}, {ID: primitive.NewObjectID(), Title: "Task 2"}}
	suite.taskUsecase.On("GetTasks", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("primitive.ObjectID")).Return(tasks, nil)

	req, _ := http.NewRequest(http.MethodGet, "/tasks", nil)
	resp := httptest.NewRecorder()
	suite.router.ServeHTTP(resp, req)

	assert.Equal(suite.T(), http.StatusOK, resp.Code)

	var receivedTasks []domain.Task
	json.NewDecoder(resp.Body).Decode(&receivedTasks)
	assert.Equal(suite.T(), tasks, receivedTasks)
	suite.taskUsecase.AssertExpectations(suite.T())
}

func (suite *TaskControllerTestSuite) TestGetTaskById() {
	taskID := primitive.NewObjectID()
	expectedTask := domain.Task{ID: taskID, Title: "Specific Task"}
	suite.taskUsecase.On("GetTaskById", mock.Anything, taskID, mock.AnythingOfType("string"), mock.AnythingOfType("primitive.ObjectID")).Return(expectedTask, nil)

	req, _ := http.NewRequest(http.MethodGet, "/tasks/"+taskID.Hex(), nil)
	resp := httptest.NewRecorder()
	suite.router.ServeHTTP(resp, req)

	assert.Equal(suite.T(), http.StatusOK, resp.Code)

	var receivedTask domain.Task
	json.NewDecoder(resp.Body).Decode(&receivedTask)
	assert.Equal(suite.T(), expectedTask, receivedTask)

	suite.taskUsecase.AssertExpectations(suite.T())
}

func (suite *TaskControllerTestSuite) TestDeleteTask() {
	taskID := primitive.NewObjectID()
	suite.taskUsecase.On("DeleteTask", mock.Anything, taskID, mock.AnythingOfType("string"), mock.AnythingOfType("primitive.ObjectID")).Return(nil)

	req, _ := http.NewRequest(http.MethodDelete, "/tasks/"+taskID.Hex(), nil)
	resp := httptest.NewRecorder()
	suite.router.ServeHTTP(resp, req)

	assert.Equal(suite.T(), http.StatusAccepted, resp.Code)
	suite.taskUsecase.AssertExpectations(suite.T())
}

func (suite *TaskControllerTestSuite) TestPutTask() {
	taskID := primitive.NewObjectID()
	task := domain.Task{ID: taskID, Title: "Updated Task", Description: "Updated Description"}
	jsonBody, _ := json.Marshal(task)
	suite.taskUsecase.On("PutTask", mock.Anything, task, taskID, mock.AnythingOfType("string"), mock.AnythingOfType("primitive.ObjectID")).Return(nil)

	req, _ := http.NewRequest(http.MethodPut, "/tasks/"+taskID.Hex(), bytes.NewBuffer(jsonBody))
	resp := httptest.NewRecorder()
	suite.router.ServeHTTP(resp, req)

	assert.Equal(suite.T(), http.StatusAccepted, resp.Code)
	suite.taskUsecase.AssertExpectations(suite.T())
}

func (suite *TaskControllerTestSuite) TestPostTask() {
	task := domain.Task{Title: "New Task", Description: "New Description"}
	jsonBody, _ := json.Marshal(task)
	suite.taskUsecase.On("PostTask", mock.Anything, task).Return(nil)

	req, _ := http.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer(jsonBody))
	resp := httptest.NewRecorder()
	suite.router.ServeHTTP(resp, req)

	assert.Equal(suite.T(), http.StatusCreated, resp.Code)
	suite.taskUsecase.AssertExpectations(suite.T())
}

func TestTaskControllerTestSuite(t *testing.T) {
	suite.Run(t, new(TaskControllerTestSuite))
}
