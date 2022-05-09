package main

import (
	"net/http"

	"github.com/callsamu/lovecalc/pkg/core"
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

	couple := core.Couple{query.Get("first"), query.Get("second")}

	if couple.FirstName == "" || couple.SecondName == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	app.render(w, r, "results.page.tmpl", &templateData{
		Match: app.calculator.Compute(couple),
	})
}
