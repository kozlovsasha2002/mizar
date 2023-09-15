package server

import (
	"database/sql"
	"github.com/go-chi/chi/v5"
	"mizar/internal/config"
	"mizar/internal/handler"
	"mizar/internal/repo"
	"mizar/internal/service"
	"net/http"
)

func NewServer(cfg *config.Config, storage *sql.DB) *http.Server {
	return &http.Server{
		Addr:         cfg.HTTPServer.Address,
		Handler:      newRouter(storage),
		ReadTimeout:  cfg.HTTPServer.Timeout,
		WriteTimeout: cfg.HTTPServer.Timeout,
		IdleTimeout:  cfg.HTTPServer.IdleTimeout,
	}
}

func newRouter(storage *sql.DB) *chi.Mux {
	router := chi.NewRouter()

	r := repo.New(storage)
	s := service.New(r)
	h := handler.New(s)

	router.Route("/events", func(r chi.Router) {
		r.Get("/", h.Event.GetAllEvents)
		r.Post("/", h.Event.Create)

		r.Route("/{eventId}", func(r chi.Router) {
			r.Use(h.Event.CtxEvent)
			r.Patch("/", h.Event.Update)
			r.Get("/", h.Event.GetById)
			r.Delete("/", h.Event.Delete)
		})
	})

	return router
}
