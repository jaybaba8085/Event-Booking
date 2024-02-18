package models

import (
	"example.com/rest-api/db"
)

// RegisteredEvent represents a registered event.
type RegisteredEvent struct {
	ID      int
	EventID int
	UserID  int
}

// RegisterUserForEvent registers a user for a specific event.
func (r *RegisteredEvent) RegisterUserForEvent(userID, eventID int) error {
	// Prepare the SQL statement
	stmt, err := db.DB.Prepare("INSERT INTO registrations (user_id, event_id) VALUES (?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute the SQL statement
	_, err = stmt.Exec(r.UserID, r.EventID)
	if err != nil {
		return err
	}

	return nil
}

// GetAllRegisteredEvents fetches all registered events from the database.
func Get_AllRegisteredEvents() ([]RegisteredEvent, error) {
	var registeredEvents []RegisteredEvent

	// Query to select all registered events
	query := `SELECT id, event_id, user_id FROM registrations`

	// Execute the query
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Iterate over the rows and scan the data into RegisteredEvent structs
	for rows.Next() {
		var registeredEvent RegisteredEvent
		err := rows.Scan(&registeredEvent.ID, &registeredEvent.EventID, &registeredEvent.UserID)
		if err != nil {
			return nil, err
		}
		registeredEvents = append(registeredEvents, registeredEvent)
	}

	// Check for errors during row iteration
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return registeredEvents, nil
}
