package usecases

import (
	"Task_Manager/domain"
	"Task_Manager/mocks/repomocks"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserUsecaseTestSuite struct {
	suite.Suite
	UserRepoMock *repomocks.UserRepository
	UserUsecase  domain.UserUsecase
}

func (s *UserUsecaseTestSuite) SetupTest() {
	s.UserRepoMock = new(repomocks.UserRepository)
	s.UserUsecase = NewUserUsecase(s.UserRepoMock, 5*time.Minute)
}

func (s *UserUsecaseTestSuite) TearDownTest() {
	s.UserRepoMock.AssertExpectations(s.T())
}

func (s *UserUsecaseTestSuite) TestRegisterUser() {
	newUser := domain.User{

		First_name: "Samrawit",
		Last_name:  "Gebermaryam",
		Password:   "sam1234",
		Email:      "samrawit@gmail.com",
		Phone:      "098900998",
		User_type:  "Admin",
	}
	s.UserRepoMock.On("RegisterUser", newUser).Return(200, nil).Once()

	status, err := s.UserUsecase.RegisterUser(context.Background(), newUser)
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), 200, status)
}

func (s *UserUsecaseTestSuite) TestLoginUser() {
	loginUser := domain.User{
		Email:    "Samrawit@gmail.com",
		Password: "12325fc",
	}
	s.UserRepoMock.On("LoginUser", loginUser).Return(200, nil, "token123").Once()

	status, err, token := s.UserUsecase.LoginUser(context.Background(), loginUser)
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), "token123", token)
	assert.Equal(s.T(), 200, status)
}

func (s *UserUsecaseTestSuite) TestDeleteUser() {
	userID := primitive.NewObjectID()
	s.UserRepoMock.On("DeleteUser", userID).Return(nil).Once()

	err := s.UserUsecase.DeleteUser(context.Background(), userID)
	assert.Nil(s.T(), err)
}

func TestUserUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(UserUsecaseTestSuite))
}
