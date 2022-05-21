package main

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRedirectLang(t *testing.T) {
	app := newTestApplication(t)
	next := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, []bytes("OK"))
	}

	cases := []struct {
		name       string
		url        string
		acceptLang string
		wantUrl    string
		wantStatus int
	}{}

	for _, ts := range cases {
		t.Run(ts.name, func(t *testing.T) {
			rr := httptest.NewRecorder()
			req, err := http.NewRequest("GET", "/en/test", nil)
			if err != nil {
				t.Fatal(err)
			}
			req.Header.Add("Accept-Language", ts.acceptLang)
			app.redirectLang(http.HandlerFunc(next)).ServeHTTP(rr, req)

			if rr.Result().Status != ts.wantStatus {
				t.Errorf("want status %d but got %d", ts.wantStatus, rr.Status)
			}

			redirect := rr.Header().Get("Localization")
			if redirect != wantUrl {
				t.Errorf("expected redirect to %s but got %s", ts.wantUrl, redirect)
			}
		})
	}
}
