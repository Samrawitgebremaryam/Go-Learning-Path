package data

import (
	"Task_Manager/models"
	"context"
	"errors"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type UserManager struct {
	client *mongo.Client
}

func NewUserManager(mClient *mongo.Client) *UserManager {
	return &UserManager{
		client: mClient,
	}
}

func (userMgr *UserManager) RegisterUser(user models.User) (int, error) {
	collection := userMgr.client.Database("taskManager").Collection("users")
	var existingUser models.User
	err := collection.FindOne(context.TODO(), bson.M{"email": user.Email}).Decode(&existingUser)
	if err == nil {
		return http.StatusConflict, errors.New("email already exists please enter different email address")
	} else if err != mongo.ErrNoDocuments {
		return http.StatusInternalServerError, err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	user.ID = primitive.NewObjectID()
	user.Password = string(hashedPassword)

	_, err = collection.InsertOne(context.TODO(), user)
	if err != nil {
		return http.StatusBadRequest, err
	}

	return http.StatusOK, nil
}

func (userMgr *UserManager) LoginUserDb(user models.User) (int, error, string) {
	collection := userMgr.client.Database("taskManager").Collection("users")

	var existingUser models.User
	err := collection.FindOne(context.TODO(), bson.M{"email": user.Email}).Decode(&existingUser)
	if err != nil {
		return http.StatusUnauthorized, errors.New("invalid email or password"), ""
	}

	err = bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password))
	if err != nil {
		return http.StatusUnauthorized, errors.New("invalid email or password"), ""
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"_id":      existingUser.ID,
		"email":    existingUser.Email,
		"usertype": existingUser.User_type,
	})

	jwtToken, err := token.SignedString([]byte("your_jwt_secret"))
	if err != nil {
		return http.StatusInternalServerError, errors.New("internal server error"), ""
	}

	return http.StatusOK, nil, jwtToken
}
