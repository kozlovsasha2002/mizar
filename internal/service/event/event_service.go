package event

import (
	"errors"
	"mizar/internal/model"
	"mizar/internal/model/dto"
	"mizar/internal/repo"
)

type ServiceEvent struct {
	r *repo.Repo
}

func New(r *repo.Repo) *ServiceEvent {
	return &ServiceEvent{r: r}
}

func (s *ServiceEvent) Save(event dto.CreationEventDto) (int, error) {
	if err := event.Validate(); err != nil {
		return 0, err
	}
	return s.r.Save(event)
}

func (s *ServiceEvent) AllEvents() ([]model.Event, error) {
	return s.r.AllEvents()
}

func (s *ServiceEvent) EventById(eventId int) (model.Event, error) {
	return s.r.EventById(eventId)
}

func (s *ServiceEvent) Update(eventId int, event dto.UpdateEventDto) error {
	return s.r.Update(eventId, event)
}

func (s *ServiceEvent) Delete(eventId int) error {
	if _, err := s.r.EventById(eventId); err != nil {
		return errors.New("event with this id does not exist")
	}
	return s.r.Delete(eventId)
}
