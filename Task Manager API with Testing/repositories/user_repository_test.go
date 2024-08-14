package repositories

import (
	"Task_Manager/domain"
	"Task_Manager/mocks/database_mocks"
	inframocks "Task_Manager/mocks/infrastructure_mocks"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// UserRepositoryTestSuite struct for holding test setup
type UserRepositoryTestSuite struct {
	suite.Suite
	repository domain.UserRepository
	mockDB     *database_mocks.UserDatabase
	mockSec    *inframocks.SecurityService
}

// SetupTest initializes before each test
func (s *UserRepositoryTestSuite) SetupTest() {
	s.mockDB = new(database_mocks.UserDatabase)
	s.mockSec = new(inframocks.SecurityService)
	s.repository = NewUserRepository(s.mockDB, s.mockSec)
}

// TestRegisterUser tests the RegisterUser method
func (s *UserRepositoryTestSuite) TestRegisterUser() {
	user := domain.User{
		Email:    "test@example.com",
		Password: "password123"}
	s.mockDB.On("FindUserByEmail", mock.Anything, user.Email).Return(domain.User{}, mongo.ErrNoDocuments).Once()
	s.mockDB.On("CreateUser", mock.Anything, mock.AnythingOfType("domain.User")).Return(nil).Once()
	s.mockSec.On("HashPassword", user.Password).Return("hashedpassword", nil).Once()

	status, err := s.repository.RegisterUser(user)

	assert.NoError(s.T(), err)
	assert.Equal(s.T(), http.StatusOK, status)
	s.mockDB.AssertExpectations(s.T())
}

// TestLoginUser tests the LoginUser method
func (s *UserRepositoryTestSuite) TestLoginUser() {
	user := domain.User{Email: "test@example.com", Password: "password123"}
	existingUser := domain.User{Email: "test@example.com", Password: "hashedpassword"}
	s.mockDB.On("FindUserByEmail", mock.Anything, user.Email).Return(existingUser, nil).Once()
	s.mockSec.On("ComparePassword", existingUser.Password, user.Password).Return(false).Once()
	s.mockSec.On("CreateToken", existingUser.ID, existingUser.Email, existingUser.User_type).Return("token", nil).Once()

	status, err, token := s.repository.LoginUser(user)

	assert.NoError(s.T(), err)
	assert.Equal(s.T(), http.StatusOK, status)
	assert.NotEmpty(s.T(), token)
	s.mockDB.AssertExpectations(s.T())
}

// TestDeleteUser tests the DeleteUser method
func (s *UserRepositoryTestSuite) TestDeleteUser() {
	id := primitive.NewObjectID()
	s.mockDB.On("DeleteUser", mock.Anything, id).Return(nil).Once()

	err := s.repository.DeleteUser(id)

	assert.NoError(s.T(), err)
	s.mockDB.AssertExpectations(s.T())
}

// TestMain is used to run the test s
func TestUserRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(UserRepositoryTestSuite))
}
