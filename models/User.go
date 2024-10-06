package models

import (
	"errors"
	"grahamkatana/api/events/db"
	"grahamkatana/api/events/utils"
)

type User struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (user *User) Save() error {
	query := `
		INSERT INTO users (name, email, password)
		VALUES (?, ?, ?)
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	result, err := stmt.Exec(user.Name, user.Email, hashedPassword)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	user.ID = uint(id)
	user.Password = hashedPassword
	return nil
}

func CheckPassword(email string, password string) error {
	query := `
		SELECT password
		FROM users
		WHERE email = ?
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	row := stmt.QueryRow(email)
	var fetchedPassword string
	err = row.Scan(&fetchedPassword)
	if err != nil {
		return err
	}
	passwordIsValid := utils.VerifyPasswordHash(password, fetchedPassword)
	if !passwordIsValid {
		return errors.New("invalid credentials")
	}
	return nil
}
