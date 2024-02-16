package models

import (
	"database/sql"
	"fmt"
	"time"

	"example.com/rest-api/db"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int
}

func (e *Event) Save() error {

	// Prepare SQL statement
	query := `
	INSERT INTO events(name, description, location, datetime, user_id) 
	VALUES(?, ?, ?, ?, ?)`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return fmt.Errorf("error preparing SQL statement: %v", err)
	}

	defer stmt.Close()

	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return fmt.Errorf("error executing SQL statement: %v", err)
	}

	// Retrieve the last inserted ID if needed
	if e.ID == 0 {
		id, err := result.LastInsertId()
		if err != nil {
			return fmt.Errorf("error retrieving last inserted ID: %v", err)
		}
		e.ID = id
	}
	return err
}

func GetAllEvents() ([]Event, error) {
	// Prepare SQL query
	query := "SELECT * FROM events"

	// Execute SQL query to retrieve rows of data
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error executing SQL query: %v", err)
	}
	defer rows.Close()

	// Create a slice to hold the events
	var events []Event

	// Iterate over the rows and scan each row into Event structs
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		if err != nil {
			return nil, fmt.Errorf("error scanning row: %v", err)
		}
		// Append the event to the slice
		events = append(events, event)
	}

	// Check for errors during iteration
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %v", err)
	}

	// Return the slice of events
	return events, nil
}

func GetEventByID(id int64) (*Event, error) {
	// Prepare SQL query
	query := "SELECT * FROM events WHERE id = ?"

	// Execute SQL query to retrieve the event with the specified ID
	row := db.DB.QueryRow(query, id)

	// Scan the row into an Event struct
	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("event with ID %d not found", id)
		}
		return nil, fmt.Errorf("error scanning row: %v", err)
	}

	// Return the event
	return &event, nil
}

func (e *Event) Update() error {
	// Check if the event already has an ID assigned
	if e.ID != 0 {
		// Prepare SQL statement for update
		query := `
            UPDATE events 
            SET name = ?, description = ?, location = ?, datetime = ?, user_id = ? 
            WHERE id = ?`

		// Execute the SQL statement
		_, err := db.DB.Exec(query, e.Name, e.Description, e.Location, e.DateTime, e.UserID, e.ID)
		if err != nil {
			return fmt.Errorf("error executing SQL statement: %v", err)
		}
		return nil
	}

	return fmt.Errorf("Event with id:  %v not found", e.ID)
}

func DeleteAllEvents() error {
	// Prepare SQL statement
	query := "DELETE FROM events"

	// Execute SQL statement to delete all events
	_, err := db.DB.Exec(query)
	if err != nil {
		return fmt.Errorf("error executing SQL statement: %v", err)
	}

	return nil
}

func DeleteEventByID(id int64) error {
	// Prepare SQL statement
	query := "DELETE FROM events WHERE id = ?"

	// Execute SQL statement to delete event by ID
	_, err := db.DB.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error executing SQL statement: %v", err)
	}

	return nil
}
