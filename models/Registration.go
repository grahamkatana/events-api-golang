package models

import "grahamkatana/api/events/db"

type Registration struct {
	ID      int `json:"id"`
	EventID int `json:"event_id"`
	UserID  int `json:"user_id"`
}

func (registration *Registration) Save() error {
	query := `
		INSERT INTO registrations (event_id, user_id)
		VALUES (?, ?)
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(registration.EventID, registration.UserID)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	registration.ID = int(id)
	return nil
}

func (registration *Registration) Delete() error {
	query := `
		DELETE FROM registrations
		WHERE id = ?
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(registration.ID)
	if err != nil {
		return err
	}
	return nil
}

func DeleteRegistration(id int, userId int) error {
	query := `
		DELETE FROM registrations
		WHERE id = ? AND user_id = ?
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id, userId)
	if err != nil {
		return err
	}
	return nil
}
