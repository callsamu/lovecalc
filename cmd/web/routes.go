package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *application) routes() http.Handler {
	r := chi.NewRouter()

	r.Use(secureHeaders, app.logRequests, app.recoverPanic)

	r.Route("/{lang:[a-z][a-z]}", func(r chi.Router) {
		r.Get("/", app.home)
		r.Get("/love", app.love)
	})

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	r.Handle("/static/", http.StripPrefix("/static", fileServer))

	return r
}
