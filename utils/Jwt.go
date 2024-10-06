package utils

import (
	"errors"
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

func VerifyJwtToken(tokenString string) (uint, error) {
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid token")
		}
		return []byte("superSecret"), nil
	})
	if err != nil {
		print(err)
		return 0, err
	}
	tokenIsValid := parsedToken.Valid
	if !tokenIsValid {
		return 0, errors.New("invalid token")
	}
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid token")
	}
	id := uint(claims["userId"].(float64))
	// email := claims["email"].(string)

	return id, nil
}
