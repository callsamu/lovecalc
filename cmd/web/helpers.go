package main

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"runtime/debug"
)

func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Println(trace)

	status := http.StatusInternalServerError
	http.Error(w, http.StatusText(status), status)
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}

func (app *application) defaultTemplateData(r *http.Request, td *templateData) {
	lang := app.lang(r)
	td.Lang = lang
	td.localizer, _ = app.localeManager.GetLocalizer(lang)
}

func (app *application) render(w http.ResponseWriter, r *http.Request, view string, td *templateData) {
	app.defaultTemplateData(r, td)
	ts, ok := app.templateCache[view]
	if !ok {
		app.serverError(w, fmt.Errorf("The template %s does not exists", view))
		return
	}

	buf := new(bytes.Buffer)
	err := ts.Execute(buf, td)
	if err != nil {
		app.serverError(w, err)
		return
	}

	_, err = buf.WriteTo(w)
	if err != nil {
		app.serverError(w, err)
	}
}

func (app *application) lang(r *http.Request) string {
	lang, ok := r.Context().Value(contextKeyLang).(string)
	if !ok {
		return ""
	}

	return lang
}

func (app *application) isLangSupported(lang string) (bool, error) {
	_, err := app.localeManager.GetLocalizer(lang)
	if errors.Is(err, errUnsupportedLocale) {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return true, nil
}
