package main

import (
	"errors"
	"net/http"

	"github.com/callsamu/lovecalc/pkg/cache"
	"github.com/callsamu/lovecalc/pkg/core"
	"github.com/callsamu/lovecalc/pkg/forms"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.RequestURI() != "/" {
		app.notFound(w)
		return
	}

	app.render(w, r, "home.page.tmpl", &templateData{})
}

func (app *application) results(w http.ResponseWriter, r *http.Request) {
	form := forms.New(r.URL.Query())
	form.Required("first", "second").
		UnicodeLettersOnly("first", "second").
		MaxLength("first", 32).
		MaxLength("second", 32)

	if !form.Valid() {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	couple := core.Couple{
		FirstName:  form.Get("first"),
		SecondName: form.Get("second"),
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
