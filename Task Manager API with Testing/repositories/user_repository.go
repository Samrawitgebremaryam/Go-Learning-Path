package repositories

import (
	"Task_Manager/domain"
	"context"
	"errors"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	userDB       domain.UserDatabase
	usersecurity domain.SecurityService
}

func NewUserRepository(userDB domain.UserDatabase, usersecurity domain.SecurityService) domain.UserRepository {
	return &UserRepository{
		userDB:       userDB,
		usersecurity: usersecurity,
	}
}

func (repo *UserRepository) RegisterUser(user domain.User) (int, error) {
	_, err := repo.userDB.FindUserByEmail(context.TODO(), user.Email)
	if err == nil {
		return http.StatusConflict, errors.New("email already exists")
	} else if err != mongo.ErrNoDocuments {
		return http.StatusInternalServerError, err
	}

	hashedPassword, err := repo.usersecurity.HashPassword(user.Password)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	user.ID = primitive.NewObjectID()
	user.Password = hashedPassword
	err = repo.userDB.CreateUser(context.TODO(), user)
	if err != nil {
		return http.StatusBadRequest, err
	}
	return http.StatusOK, nil
}

func (repo *UserRepository) LoginUser(user domain.User) (int, error, string) {
	existingUser, err := repo.userDB.FindUserByEmail(context.TODO(), user.Email)
	if err != nil {
		return http.StatusUnauthorized, errors.New("invalid email or password"), ""
	}

	if pass := repo.usersecurity.ComparePassword(existingUser.Password, user.Password); pass {
		return http.StatusUnauthorized, errors.New("invalid email or password"), ""
	}

	token, err := repo.usersecurity.CreateToken(existingUser.ID, existingUser.Email, existingUser.User_type)
	if err != nil {
		return http.StatusInternalServerError, err, ""
	}
	return http.StatusOK, nil, token
}

func (repo *UserRepository) DeleteUser(id primitive.ObjectID) error {
	return repo.userDB.DeleteUser(context.TODO(), id)
}
