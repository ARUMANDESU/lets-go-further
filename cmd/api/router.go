package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) routes() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.NotFound(app.notFoundResponse)
	r.MethodNotAllowed(app.methodNotAllowedResponse)

	r.Route("/v1", func(r chi.Router) {
		r.Get("/healthcheck", app.healthcheckHandler)

		r.Route("/movies", func(r chi.Router) {
			r.Post("/", app.createMovieHandler)
			r.Get("/{id}", app.showMovieHandler)
			r.Put("/{id}", app.updateMovieHandler)

		})
	})

	return r
}
