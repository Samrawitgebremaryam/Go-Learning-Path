package infrastructure

import (
	"Task_Manager/domain"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(user domain.User) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	return string(hashedPassword), err
}

func ComparePassword(existingUser domain.User, user domain.User) error {
	return bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password))
}
