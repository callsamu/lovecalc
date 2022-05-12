package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/callsamu/lovecalc/pkg/core"
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

func TestResults(t *testing.T) {
	app := newTestApplication(t)

	t.Run("test sample computation", func(t *testing.T) {
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
	})

	t.Run("test not supplied parameters", func(t *testing.T) {
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/results", nil)
		app.results(rr, req)

		if rr.Code != http.StatusSeeOther {
			t.Errorf("expected status %d, but got %d", http.StatusSeeOther, rr.Code)
		}
	})
}
