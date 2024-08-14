package infrastructure

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateToken generates a JWT token for a given user
func (s *SecurityService) CreateToken(id primitive.ObjectID, email string, user_type string) (string, error) {
	claims := jwt.MapClaims{
		"sub":       id.Hex(),
		"email":     email,
		"user_type": user_type,
		"exp":       time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secretKey := "your_jwt_secret" // This should be stored securely
	return token.SignedString([]byte(secretKey))
}

// ValidateToken checks if the provided token is valid and not expired
func (s *SecurityService) ValidateToken(tokenstr string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenstr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("your_jwt_secret"), nil
	})
	return token, err
}
