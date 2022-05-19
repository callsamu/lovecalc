package main

import (
	"html/template"
	"math"
	"path/filepath"

	"github.com/callsamu/lovecalc/pkg/core"
	"github.com/callsamu/lovecalc/pkg/forms"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type templateData struct {
	Lang      string
	Match     *core.Match
	Form      *forms.Form
	Localizer *i18n.Localizer
}

func toRoundedPercentage(x float64) float64 {
	return math.Round(10000*x) / 100
}

// Translate function
func (td *templateData) t(key string) (string, error) {
	return td.Localizer.Localize(&i18n.LocalizeConfig{
		MessageID:    key,
		TemplateData: td,
	})
}

var functions = template.FuncMap{
	"toRoundedPercentage": toRoundedPercentage,
}

func newTemplateCache(dir string) (map[string]*template.Template, error) {
	templateCache := map[string]*template.Template{}

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

		templateCache[name] = ts
	}

	return templateCache, nil
}
