package handler

import (
	"mizar/internal/handler/event"
	"mizar/internal/service"
)

type Handler struct {
	Event *event.HandlerEvent
}

func New(s *service.Service) *Handler {
	return &Handler{
		Event: event.New(s),
	}
}
