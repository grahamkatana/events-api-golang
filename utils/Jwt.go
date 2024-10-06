package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJwtToken(email string, userId uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 3).Unix(),
	})
	tokenString, err := token.SignedString([]byte("superSecret"))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
