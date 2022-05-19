package forms

import (
	"testing"

	"github.com/callsamu/lovecalc/pkg/translations"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

func TestErrors(t *testing.T) {
	lang := language.MustParse("en")
	bundle, err := translations.Load(translations.LocalesFS, lang)
	if err != nil {
		t.Fatal(err)
	}

	l := i18n.NewLocalizer(bundle, "en")
	e := errors{localizer: l}
	e.Add("foo", "FooError")

	want := "foo"
	got, err := e.Get("foo")
	if err != nil {
		t.Fatal(err)
	}

	if want != got {
		t.Errorf("want %s; got %s on error map", want, got)
	}

	e.Add("foo", "BarError")
	got, err = e.Get("foo")
	if err != nil {
		t.Fatal(err)
	}

	if want != got {
		t.Errorf("want %s; got %s on error map", want, got)
	}
}
