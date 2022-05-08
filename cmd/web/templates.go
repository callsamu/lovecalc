package main

import (
	"html/template"
	"math"
	"path/filepath"
)

type templateData struct {
	FirstName  string
	SecondName string
	Chance     float64
	MixedName  string
}

func toRoundedPercentage(x float64) float64 {
	return math.Round(10000*x) / 100
}

var functions = template.FuncMap{
	"toRoundedPercentage": toRoundedPercentage,
}

func newTemplateCache(dir string) (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob(filepath.Join(dir, "*.page.tmpl"))
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob(filepath.Join(dir, "*.layout.tmpl"))
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return cache, nil
}
