package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRedirectLang(t *testing.T) {
	app := newTestApplication(t)
	next := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, []byte("OK"))
	}

	cases := []struct {
		name       string
		url        string
		acceptLang string
		wantUrl    string
		wantStatus int
	}{
		{
			name:       "simple case",
			url:        "/en/test",
			acceptLang: "pt",
			wantUrl:    "/pt/test",
			wantStatus: http.StatusSeeOther,
		},
	}

	for _, ts := range cases {
		t.Run(ts.name, func(t *testing.T) {
			rr := httptest.NewRecorder()
			req, err := http.NewRequest("GET", "/en/test", nil)
			if err != nil {
				t.Fatal(err)
			}
			req.Header.Add("Accept-Language", ts.acceptLang)
			app.redirectLang(http.HandlerFunc(next)).ServeHTTP(rr, req)

			status := rr.Result().StatusCode
			if status != ts.wantStatus {
				t.Errorf("want status %d but got %d", ts.wantStatus, status)
			}

			t.Log(rr.Header())

			redirect := rr.Header().Get("Location")
			if redirect == "" {
				t.Errorf("redirect URL is empty")
				return
			}
			if redirect != ts.wantUrl {
				t.Errorf("expected redirect to %s but got %s", ts.wantUrl, redirect)
			}
		})
	}
}
