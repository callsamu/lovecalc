package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/callsamu/lovecalc/pkg/cache/mock"
	"github.com/callsamu/lovecalc/pkg/core"
	"github.com/callsamu/lovecalc/pkg/translations"
	"golang.org/x/text/language"
)

type testServer struct {
	*httptest.Server
}

func newTestApplication(t *testing.T) *application {
	c, err := core.NewCalculator(core.HashWithFNV)
	if err != nil {
		t.Fatal(err)
	}

	infoLog := log.New(ioutil.Discard, "", 0)
	errorLog := log.New(os.Stderr, "TEST ERROR:\t", 0)

	if err != nil {
		t.Fatal(err)
	}

	bundle, err := translations.Load(translations.LocalesFS, language.English)
	if err != nil {
		t.Fatal(err)
	}

	app := &application{
		calculator:    c,
		infoLog:       infoLog,
		errorLog:      errorLog,
		matchCache:    mock.NewMatchCache(),
		localizers:    newLocalizers(bundle),
		templateCache: map[string]*template.Template{},
	}

	app.initTemplateCache("./../../ui/template/")
	return app
}

func newTestServer(t *testing.T, h http.Handler) *testServer {
	ts := httptest.NewServer(h)
	return &testServer{ts}
}

func (ts *testServer) get(t *testing.T, uri string) (int, http.Header, []byte) {
	resp, err := ts.Client().Get(ts.URL + uri)
	if err != nil {
		t.Error(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}

	return resp.StatusCode, resp.Header, body
}
