package models

import (
	"grahamkatana/api/events/db"
	"time"
)

type Event struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Date        time.Time `json:"date" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	UserID      uint      `json:"user_id"`
}

var events []Event = []Event{}

func (event *Event) Save() error {
	query := `
	INSERT INTO events (title, description, date, location, user_id)
	VALUES (?, ?, ?, ?, ?)
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(event.Title, event.Description, event.Date, event.Location, event.UserID)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	event.ID = uint(id)
	return nil

}

func GetAllEvents() ([]Event, error) {
	rows, err := db.DB.Query("SELECT * FROM events")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Title, &event.Description, &event.Date, &event.Location, &event.UserID)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}

func GetEventByID(id uint) (*Event, error) {
	var event Event
	row := db.DB.QueryRow("SELECT * FROM events WHERE id = ?", id)
	err := row.Scan(&event.ID, &event.Title, &event.Description, &event.Date, &event.Location, &event.UserID)
	if err != nil {
		return nil, err
	}
	return &event, nil
}

func (event Event) UpdateEvent() error {
	query := `
	UPDATE events
	SET title = ?, description = ?, date = ?, location = ?
	WHERE id = ?
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(event.Title, event.Description, event.Date, event.Location, event.ID)
	if err != nil {
		return err
	}
	return nil
}

func (event Event) DeleteEvent() error {
	query := `
	DELETE FROM events
	WHERE id = ?
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(event.ID)
	if err != nil {
		return err
	}
	return nil

}
