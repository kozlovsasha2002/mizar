package repo

import (
	"database/sql"
	"mizar/internal/model"
	"mizar/internal/model/dto"
	"mizar/internal/repo/event"
)

type Repo struct {
	Event
}

func New(db *sql.DB) *Repo {
	return &Repo{
		Event: event.New(db),
	}
}

type Event interface {
	Save(event dto.CreationEventDto) (int, error)
	AllEvents() ([]model.Event, error)
	EventById(eventId int) (model.Event, error)
	Update(eventId int, eventDto dto.UpdateEventDto) error
	Delete(eventId int) error
}
