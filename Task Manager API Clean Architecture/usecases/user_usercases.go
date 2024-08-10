package usecases

import (
	"Task_Manager/domain"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserUsecase struct {
	userRepo domain.UserRepository
	timeout  time.Duration
}

func NewUserUsecase(userRepo domain.UserRepository, timeout time.Duration) domain.UserUsecase {
	return &UserUsecase{
		userRepo: userRepo,
		timeout:  timeout,
	}
}

func (u *UserUsecase) RegisterUser(ctx context.Context, user domain.User) (int, error) {
	_, cancel := context.WithTimeout(ctx, u.timeout)
	defer cancel()

	return u.userRepo.RegisterUser(user)
}

func (u *UserUsecase) LoginUser(ctx context.Context, user domain.User) (int, error, string) {
	_, cancel := context.WithTimeout(ctx, u.timeout)
	defer cancel()

	return u.userRepo.LoginUser(user)

}

func (u *UserUsecase) DeleteUser(ctx context.Context, id primitive.ObjectID) error {
	_, cancel := context.WithTimeout(ctx, u.timeout)
	defer cancel()

	return u.userRepo.DeleteUser(id)

}
