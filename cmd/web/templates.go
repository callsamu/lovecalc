package main

import (
	"fmt"
	"html/template"
	"path/filepath"
)

type templateData struct {
	Result float64
}

func toRoundedPercentage(x float64) string {
	return fmt.Sprintf("%.2f%%", x*100)
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
