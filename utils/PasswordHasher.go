package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func VerifyPasswordHash(plainPassword, hashedPassword string) bool {
	error := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	return error == nil
}
