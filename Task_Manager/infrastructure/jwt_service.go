package infrastructure

import (
	"Task_Manager/domain"
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

func CreateToken(existingUser domain.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"_id":      existingUser.ID,
		"email":    existingUser.Email,
		"usertype": existingUser.User_type,
	})

	jwtToken, err := token.SignedString([]byte("your_jwt_secret"))
	return jwtToken, err

}

func ValidateToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte("your_jwt_secret"), nil
	})
	return token, err

}
