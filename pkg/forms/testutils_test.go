package forms

import (
	"io/fs"
	"testing"
	"testing/fstest"
	"time"

	"github.com/callsamu/lovecalc/pkg/translations"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

const errorMessages = `
	[FooError]
	other = "foo"
	[BarError]
	other = "bar"
`

func errMsgBundle(t *testing.T) *i18n.Bundle {
	fs := fstest.MapFS{
		"mock.en.toml": &fstest.MapFile{
			Data:    []byte(errorMessages),
			Mode:    fs.ModeType,
			ModTime: time.Now(),
			Sys:     nil,
		},
	}

	bundle, err := translations.Load(fs, language.English)
	if err != nil {
		t.Fatal(err)
	}

	return bundle
}

func newLocalizer(t *testing.T) *i18n.Localizer {
	bundle := errMsgBundle(t)
	return i18n.NewLocalizer(bundle, "en")
}
