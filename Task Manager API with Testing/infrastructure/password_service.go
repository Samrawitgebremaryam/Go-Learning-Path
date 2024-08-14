package infrastructure

import (
	"Task_Manager/domain"

	"golang.org/x/crypto/bcrypt"
)

type SecurityService struct{}

func NewSecurityService() domain.SecurityService {
	return &SecurityService{}
}

// HashPassword hashes the provided password using bcrypt
func (s *SecurityService) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword), err
}

// ComparePassword compares the provided password with the stored hash
func (s *SecurityService) ComparePassword(hash string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
