package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
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

func TestLoveE2E(t *testing.T) {
	app := newTestApplication(t)

	srv := newTestServer(t, app.routes())
	match := app.calculator.Compute(core.Couple{
		FirstName:  "ha",
		SecondName: "he",
	})
	prob := fmt.Sprintf("%d", int(toRoundedPercentage(match.Probability)))

	status, _, body := srv.get(t, "/en/love?first=ha&second=he")

	if status != http.StatusOK {
		t.Errorf("expected status 200 OK, but got %d", status)
	}

	if !bytes.Contains(body, []byte(prob)) {
		t.Errorf("expected body to contain substring %s", prob)
	}
}

func TestLove(t *testing.T) {
	app := newTestApplication(t)

	cases := []struct {
		name   string
		first  string
		second string
	}{
		{"test not supplied parameters", "", ""},
		{"test invalid parameters", "!!$", "!$#"},
	}

	for _, ts := range cases {
		t.Run(ts.name, func(t *testing.T) {
			rr := httptest.NewRecorder()
			url := fmt.Sprintf("/en/love?first=%s&second=%s", ts.first, ts.second)
			req, _ := http.NewRequest("GET", url, nil)

			form := forms.New(req.URL.Query())
			form.Required("first", "second").
				UnicodeLettersOnly("first", "second").
				MaxLength("first", 32).
				MaxLength("second", 32)

			app.love(rr, req)
			body, err := ioutil.ReadAll(rr.Result().Body)
			if err != nil {
				t.Error(err)
				return
			}

			firstErr := []byte(form.Errors.Get("first"))
			secondErr := []byte(form.Errors.Get("second"))

			if !bytes.Contains(body, firstErr) {
				t.Errorf("expected body to contain substring \"%s\"", firstErr)
			}
			if !bytes.Contains(body, secondErr) {
				t.Errorf("expected body to contain substring \"%s\"", secondErr)
			}
		})
	}
}
