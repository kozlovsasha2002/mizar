package event

import (
	"database/sql"
	"fmt"
	"mizar/internal/model"
	"mizar/internal/model/dto"
	"strings"
)

const (
	eventsTable = "events"
)

type RepoEvent struct {
	db *sql.DB
}

func New(db *sql.DB) *RepoEvent {
	return &RepoEvent{db: db}
}

func (r *RepoEvent) Save(event dto.CreationEventDto) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s "+
		"(link, deadline_date, description) "+
		"VALUES ($1, $2, $3) RETURNING event_id", eventsTable)

	row := r.db.QueryRow(query, event.Link, event.DeadlineDate, event.Description)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *RepoEvent) AllEvents() ([]model.Event, error) {
	events := make([]model.Event, 0)
	query := fmt.Sprintf("SELECT * FROM %s", eventsTable)

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var ev model.Event
		err := rows.Scan(&ev.EventId, &ev.Link, &ev.DeadlineDate, &ev.IsCompleted, &ev.Description)
		if err != nil {
			return nil, err
		}
		events = append(events, ev)
	}

	return events, nil
}

func (r *RepoEvent) EventById(eventId int) (model.Event, error) {
	var event model.Event
	query := fmt.Sprintf("SELECT * FROM %s WHERE event_id = $1", eventsTable)

	row := r.db.QueryRow(query, eventId)
	if err := row.Scan(&event.EventId, &event.Link, &event.DeadlineDate, &event.IsCompleted, &event.Description); err != nil {
		return event, err
	}

	return event, nil
}

func (r *RepoEvent) Update(eventId int, event dto.UpdateEventDto) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if event.Link != nil {
		setValues = append(setValues, fmt.Sprintf("link=$%d", argId))
		args = append(args, *event.Link)
		argId++
	}
	if event.DeadlineDate != nil {
		setValues = append(setValues, fmt.Sprintf("deadline_date=$%d", argId))
		args = append(args, *event.DeadlineDate)
		argId++
	}
	if event.IsCompleted != nil {
		setValues = append(setValues, fmt.Sprintf("is_completed=$%d", argId))
		args = append(args, *event.IsCompleted)
		argId++
	}
	if event.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *event.Description)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s SET %s WHERE event_id = $%d",
		eventsTable, setQuery, argId)
	args = append(args, eventId)

	_, err := r.db.Exec(query, args...)
	return err
}

func (r *RepoEvent) Delete(eventId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE event_id = $1", eventsTable)
	_, err := r.db.Exec(query, eventId)
	return err
}
