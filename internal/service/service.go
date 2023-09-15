package service

import (
	"mizar/internal/repo"
	"mizar/internal/service/event"
)

type Service struct {
	Event *event.ServiceEvent
}

func New(r *repo.Repo) *Service {
	return &Service{
		Event: event.New(r),
	}
}
