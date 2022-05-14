package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/callsamu/lovecalc/pkg/core"
	"github.com/callsamu/lovecalc/pkg/forms"
)

func TestHome(t *testing.T) {
	app := newTestApplication(t)
	req, _ := http.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()

	app.home(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected status 200 OK, but got %d", rr.Code)
	}
}

func TestResultsE2E(t *testing.T) {
	app := newTestApplication(t)

	srv := newTestServer(t, app.routes())
	match := app.calculator.Compute(core.Couple{
		FirstName:  "ha",
		SecondName: "he",
	})
	prob := fmt.Sprintf("%d", int(toRoundedPercentage(match.Probability)))

	status, _, body := srv.get(t, "/results?first=ha&second=he")

	if status != http.StatusOK {
		t.Errorf("expected status 200 OK, but got %d", status)
	}

	if !bytes.Contains(body, []byte(prob)) {
		t.Errorf("expected body to contain substring %s", prob)
	}
}

func TestResults(t *testing.T) {
	app := newTestApplication(t)

	cases := []struct {
		name       string
		first      string
		second     string
		wantStatus int
	}{
		{"test not supplied parameters", "", "", http.StatusSeeOther},
		{"test invalid parameters", "!!$", "!$#", http.StatusSeeOther},
	}

	for _, ts := range cases {
		t.Run(ts.name, func(t *testing.T) {
			rr := httptest.NewRecorder()
			url := fmt.Sprintf("/results?first=%s&second=%s", ts.first, ts.second)
			req, _ := http.NewRequest("GET", url, nil)
			app.results(rr, req)

			if rr.Code != ts.wantStatus {
				t.Errorf("expected status %d, but got %d", ts.wantStatus, rr.Code)
			}

			form := req.Context().Value(formCtxKey)
			if form == nil {
				t.Error("expected request context to contain a form")
			}
		})
	}
}
