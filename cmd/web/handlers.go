package main

import (
	"errors"
	"net/http"

	"github.com/callsamu/lovecalc/pkg/cache"
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

	couple := core.Couple{
		FirstName:  query.Get("first"),
		SecondName: query.Get("second"),
	}

	if couple.FirstName == "" || couple.SecondName == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	match, err := app.matchCache.Get(couple)
	if err != nil {
		switch {
		case errors.Is(err, cache.ErrKeyNotFound):
			match = app.calculator.Compute(couple)
			err := app.matchCache.Set(couple, match)
			if err != nil {
				app.serverError(w, err)
				return
			}
		default:
			app.serverError(w, err)
			return
		}
	}

	app.render(w, r, "results.page.tmpl", &templateData{
		Match: match,
	})
}
