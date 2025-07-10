package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
	"time"
)

type application struct {
	config config
}

type config struct {
	addr string
}

func (app *application) mount() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)
	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/health", app.healthHandler)
	})

	return r
}

func (app *application) run(mux *chi.Mux) error {
	srv := &http.Server{
		Addr:              app.config.addr,
		Handler:           mux,
		WriteTimeout:      time.Second * 30,
		ReadHeaderTimeout: time.Second * 10,
		IdleTimeout:       time.Minute,
	}

	log.Printf("Server has started on port %s", app.config.addr)

	return srv.ListenAndServe()
}
