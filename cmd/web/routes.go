package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *application) routes() http.Handler {
	r := chi.NewRouter()

	r.Use(secureHeaders)
	r.Use(app.logRequests)
	r.Use(app.recoverPanic)

	r.Get("/", app.home)
	r.Get("/results", app.results)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	r.Handle("/static/", http.StripPrefix("/static", fileServer))

	return r
}
