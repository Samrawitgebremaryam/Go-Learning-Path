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

// UserControllerTestSuite struct to hold necessary objects for testing the UserController.
type UserControllerTestSuite struct {
	suite.Suite
	router      *gin.Engine
	userUsecase *usecase_mocks.UserUsecase
}

// SetupTest initializes testing environment and mocks before each test.
func (suite *UserControllerTestSuite) SetupTest() {
	gin.SetMode(gin.TestMode)
	suite.router = gin.Default()
	suite.userUsecase = new(usecase_mocks.UserUsecase)
	userController := NewUserController(suite.userUsecase)
	suite.router.POST("/register", userController.RegisterUser)
	suite.router.POST("/login", userController.LoginUser)
	suite.router.DELETE("/users/:id", userController.DeleteUser)
}

// TestRegisterUser tests the RegisterUser function of the UserController.
func (suite *UserControllerTestSuite) TestRegisterUser() {
	user := domain.User{Email: "test@example.com", Password: "password123"}
	userJSON, _ := json.Marshal(user)
	suite.userUsecase.On("RegisterUser", mock.Anything, user).Return(http.StatusOK, nil)

	req, _ := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(userJSON))
	resp := httptest.NewRecorder()
	suite.router.ServeHTTP(resp, req)

	assert.Equal(suite.T(), http.StatusOK, resp.Code)
	suite.userUsecase.AssertExpectations(suite.T())
}

// TestLoginUser tests the LoginUser function of the UserController.
func (suite *UserControllerTestSuite) TestLoginUser() {
	user := domain.User{Email: "test@example.com", Password: "password123"}
	userJSON, _ := json.Marshal(user)
	suite.userUsecase.On("LoginUser", mock.Anything, user).Return(http.StatusOK, nil, "token123")

	req, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(userJSON))
	resp := httptest.NewRecorder()
	suite.router.ServeHTTP(resp, req)

	assert.Equal(suite.T(), http.StatusOK, resp.Code)
	assert.Contains(suite.T(), resp.Body.String(), "token123")
	suite.userUsecase.AssertExpectations(suite.T())
}

// TestDeleteUser tests the DeleteUser function of the UserController.
func (suite *UserControllerTestSuite) TestDeleteUser() {
	userID := primitive.NewObjectID()
	suite.userUsecase.On("DeleteUser", mock.Anything, userID).Return(nil)

	req, _ := http.NewRequest(http.MethodDelete, "/users/"+userID.Hex(), nil)
	resp := httptest.NewRecorder()
	suite.router.ServeHTTP(resp, req)

	assert.Equal(suite.T(), http.StatusAccepted, resp.Code)
	suite.userUsecase.AssertExpectations(suite.T())
}

// Run the suite
func TestUserController(t *testing.T) {
	suite.Run(t, new(UserControllerTestSuite))
}
