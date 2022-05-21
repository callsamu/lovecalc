package main

import (
	"context"
	"net/http"
	"testing"
)

func TestLangRetrieve(t *testing.T) {
	app := newTestApplication(t)
	req, _ := http.NewRequest("GET", "/foo", nil)
	ctx := context.WithValue(req.Context(), contextKeyLang, "en")
	lang := app.lang(req.WithContext(ctx))

	if lang != "en" {
		t.Errorf("want en; got %s", lang)
	}
}
