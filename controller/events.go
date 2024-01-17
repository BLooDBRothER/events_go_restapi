package controller

import (
	"example.com/events/db"
	"example.com/events/models"
)

func GetAllEvents() ([]models.Event, error) {
	var events []models.Event
	query := "SELECT * FROM events"

	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var event models.Event
		rows.Scan(&event.Id, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)
		events = append(events, event)
	}

	return events, nil
}

func GetEventById(id int64) (*models.Event, error) {
	query := "SELECT * FROM events WHERE id=?"

	row := db.DB.QueryRow(query, id)

	var event models.Event
	err := row.Scan(&event.Id, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)

	if err != nil {
		return nil, err
	}

	return &event, nil
}

func CreateEvent(event *models.Event) (int64, error) {
	query := `
	INSERT INTO events(name, description, location, date_time, user_id) 
	VALUES (?, ?, ?, ?, ?)`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	result, err := stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.UserId)

	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

func UpdateEvent(event *models.Event) error {
	query := `
	UPDATE events
	SET name = ?, description = ?, location = ?, date_time = ?
	WHERE id = ?`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.Id)

	if err != nil {
		return err
	}
	return nil
}


func DeleteEvent(eventId int64) error {
	query := `DELETE FROM events WHERE id = ?`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(eventId)

	if err != nil {
		return err
	}
	return nil
}
