package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
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
		result := app.calculator.Compute("ha", "he")
		resultString := fmt.Sprintf("%f", result)

		status, _, body := srv.get(t, "/results?first=ha&second=he")

		if status != http.StatusOK {
			t.Errorf("expected status 200 OK, but got %d", status)
		}

		if !bytes.Contains(body, []byte(resultString)) {
			t.Errorf("expected body to contain substring %s", resultString)
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
