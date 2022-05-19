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

var functions = template.FuncMap{
	"toRoundedPercentage": toRoundedPercentage,
}

func newTemplateCache(dir string, lm *LocaleManager) (map[string]*template.Template, error) {
	templateCache := map[string]*template.Template{}

	pages, err := filepath.Glob(filepath.Join(dir, "*.page.tmpl"))
	if err != nil {
		return nil, err
	}

	funcs := functions
	funcs["t"] = lm.localize // Set translator function

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).Funcs(funcs).ParseFiles(page)
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob(filepath.Join(dir, "*.layout.tmpl"))
		if err != nil {
			return nil, err
		}

		templateCache[name] = ts
	}

	return templateCache, nil
}
