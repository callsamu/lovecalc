package main

import (
	"net/http"

	mixer "github.com/callsamu/lovecalc/pkg/namemixer"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.RequestURI() != "/" {
		app.notFound(w)
		return
	}

	app.render(w, r, "home.page.tmpl", &templateData{})
}

func (app *application) results(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	first := query.Get("first")
	second := query.Get("second")

	if first == "" || second == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	result := app.calculator.Compute(first, second)
	name := mixer.MixNames(first, second)

	app.render(w, r, "results.page.tmpl", &templateData{
		Chance:     result,
		MixedName:  name,
		FirstName:  first,
		SecondName: second,
	})
}
