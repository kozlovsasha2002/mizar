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

// Create godoc
// @Summary Create
// @Description create new event
// @Tags events
// @Accept json
// @Produce json
// @Param DataToCreate body dto.CreationEventDto true "json with data about a new event"
// @Success 200 {object} resp.Response
// @Failure 400 {object} er.ErrResponse
// @Failure 500 {object} er.ErrResponse
// @Router /events [post]
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

// GetAllEvents godoc
// @Summary GetAllEvents
// @Description get all events
// @Tags events
// @Accept json
// @Produce json
// @Success 200 {array} resp.Response
// @Failure 400 {object} er.ErrResponse
// @Failure 422 {object} er.ErrResponse
// @Router /events [get]
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

// GetById godoc
// @Summary GetById
// @Description get event by id
// @Tags events
// @Accept json
// @Produce json
// @Param eventId	path	int		true	"Event id"
// @Success 200 {object} resp.Response
// @Failure 400 {object} er.ErrResponse
// @Failure 404 {object} er.ErrResponse
// @Router /events/{eventId} [get]
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

// Update godoc
// @Summary Update
// @Description update event by id
// @Tags events
// @Accept json
// @Produce json
// @Param eventId	path	int		true	"Event id"
// @Param DataToUpdate body dto.UpdateEventDto true "json with data about a new event"
// @Success 200 {object} resp.Response
// @Failure 400 {object} er.ErrResponse
// @Failure 422 {object} er.ErrResponse
// @Failure 500 {object} er.ErrResponse
// @Router /events/{eventId} [patch]
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

// Delete godoc
// @Summary Delete
// @Description delete event by id
// @Tags events
// @Accept json
// @Produce json
// @Param eventId	path	int		true	"Event id"
// @Success 200 {object} resp.Response
// @Failure 400 {object} er.ErrResponse
// @Failure 404 {object} er.ErrResponse
// @Failure 422 {object} er.ErrResponse
// @Router /events/{eventId} [delete]
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
