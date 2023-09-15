package event

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"mizar/internal/handler/er"
	"mizar/internal/handler/resp"
	"mizar/internal/model/dto"
	"mizar/internal/service"
	"net/http"
	"strconv"
)

type HandlerEvent struct {
	s *service.Service
}

func New(s *service.Service) *HandlerEvent {
	return &HandlerEvent{s: s}
}

func (e *HandlerEvent) CtxEvent(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		eventIdStr := chi.URLParam(r, "eventId")

		eventId, err := strconv.Atoi(eventIdStr)
		if err != nil {
			render.JSON(w, r, er.ErrInvalidRequest(err))
			return
		}

		ctx := context.WithValue(r.Context(), "eventId", eventId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (e *HandlerEvent) Create(w http.ResponseWriter, r *http.Request) {
	var event dto.CreationEventDto
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		render.JSON(w, r, er.ErrInvalidRequest(err))
		return
	}
	eventId, err := e.s.Event.Save(event)
	if err != nil {
		render.JSON(w, r, er.ErrInternal(err))
		return
	}
	render.JSON(w, r, resp.ResponseWithId{
		HTTPStatusCode: http.StatusCreated,
		Id:             eventId,
	})
}

func (e *HandlerEvent) GetAllEvents(w http.ResponseWriter, r *http.Request) {
	events, err := e.s.Event.AllEvents()
	if err != nil {
		render.JSON(w, r, er.ErrUnprocessableEntity(err))
		return
	}
	render.JSON(w, r, resp.Response{
		HTTPStatusCode: http.StatusOK,
		Value:          events,
	})
}

func (e *HandlerEvent) GetById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	eventId, ok := ctx.Value("eventId").(int)
	if !ok {
		render.JSON(w, r, er.ErrInvalidRequest(errors.New("eventId param is not correct")))
		return
	}
	event, err := e.s.Event.EventById(eventId)
	if err != nil {
		render.JSON(w, r, er.ErrNotFound(err))
		return
	}
	render.JSON(w, r, resp.Response{
		HTTPStatusCode: http.StatusOK,
		Value:          event,
	})
}

func (e *HandlerEvent) Update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	eventId, ok := ctx.Value("eventId").(int)
	if !ok {
		render.JSON(w, r, er.ErrUnprocessableEntity(errors.New("eventId param is not correct")))
		return
	}
	var eventDto dto.UpdateEventDto
	err := json.NewDecoder(r.Body).Decode(&eventDto)
	if err != nil {
		render.JSON(w, r, er.ErrInvalidRequest(err))
		return
	}
	err = e.s.Event.Update(eventId, eventDto)
	if err != nil {
		render.JSON(w, r, er.ErrInternal(err))
		return
	}
	render.JSON(w, r, resp.Response{
		HTTPStatusCode: http.StatusOK,
		Value:          "OK",
	})
}

func (e *HandlerEvent) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	eventId, ok := ctx.Value("eventId").(int)
	if !ok {
		render.JSON(w, r, er.ErrUnprocessableEntity(errors.New("eventId param is not correct")))
		return
	}
	err := e.s.Event.Delete(eventId)
	if err != nil {
		render.JSON(w, r, er.ErrNotFound(err))
		return
	}
	render.JSON(w, r, resp.Response{
		HTTPStatusCode: http.StatusOK,
		Value:          "OK",
	})
}
