package repositories

import (
	"Task_Manager/domain"
	"Task_Manager/infrastructure"

	"context"
	"errors"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewUserRepository(mclient *mongo.Client) domain.UserRepository {
	collection := mclient.Database("taskManager").Collection("users")
	return &UserRepository{
		client:     mclient,
		collection: collection,
	}
}

func (userMgr *UserRepository) RegisterUser(user domain.User) (int, error) {
	collection := userMgr.client.Database("taskManager").Collection("users")
	var existingUser domain.User
	err := collection.FindOne(context.TODO(), bson.M{"email": user.Email}).Decode(&existingUser)
	if err == nil {
		return http.StatusConflict, errors.New("email already exists please enter different email address")
	} else if err != mongo.ErrNoDocuments {
		return http.StatusInternalServerError, err
	}

	hashedPassword, err := infrastructure.HashPassword(user)

	if err != nil {
		return http.StatusInternalServerError, err
	}

	user.ID = primitive.NewObjectID()
	user.Password = hashedPassword

	_, err = collection.InsertOne(context.TODO(), user)
	if err != nil {
		return http.StatusBadRequest, err
	}

	return http.StatusOK, nil
}

func (userMgr *UserRepository) LoginUser(user domain.User) (int, error, string) {
	collection := userMgr.client.Database("taskManager").Collection("users")

	var existingUser domain.User
	err := collection.FindOne(context.TODO(), bson.M{"email": user.Email}).Decode(&existingUser)
	if err != nil {
		return http.StatusUnauthorized, errors.New("invalid email or password"), ""
	}

	err = infrastructure.ComparePassword(existingUser, user)
	if err != nil {
		return http.StatusUnauthorized, errors.New("invalid email or password"), ""
	}

	jwtToken, err := infrastructure.CreateToken(existingUser)
	if err != nil {
		return http.StatusInternalServerError, errors.New("internal server error"), ""
	}

	return http.StatusOK, nil, jwtToken
}

// DeleteUser deletes a specific task by its ID
func (userMgr *UserRepository) DeleteUser(id primitive.ObjectID) error {
	filter := bson.D{{Key: "_id", Value: id}}

	result, err := userMgr.collection.DeleteOne(context.TODO(), filter)

	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("user not found invalid ID please input another ID")
	}
	return nil
}
