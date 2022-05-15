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

	r.Get("/{lang}/", app.home)
	r.Get("/{lang}/love", app.love)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	r.Handle("/static/", http.StripPrefix("/static", fileServer))

	return r
}
