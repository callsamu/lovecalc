package main

import (
	"html/template"
	"math"
	"path/filepath"

	"github.com/callsamu/lovecalc/pkg/core"
	"github.com/callsamu/lovecalc/pkg/forms"
)

type templateData struct {
	Lang  string
	Match *core.Match
	Form  *forms.Form
}

func toRoundedPercentage(x float64) float64 {
	return math.Round(10000*x) / 100
}

func (app *application) initTemplateCache(dir string) error {
	pages, err := filepath.Glob(filepath.Join(dir, "*.page.tmpl"))
	if err != nil {
		return err
	}

	functions := template.FuncMap{
		"t":                   newLocalizerFunc(app.localizers),
		"toRoundedPercentage": toRoundedPercentage,
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return err
		}

		ts, err = ts.ParseGlob(filepath.Join(dir, "*.layout.tmpl"))
		if err != nil {
			return err
		}

		app.templateCache[name] = ts
	}

	return nil
}
