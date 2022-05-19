package translations

import (
	"io/fs"
	"testing"
	"testing/fstest"
	"time"

	"golang.org/x/text/language"
)

const mockTranslation = `
	[foo]
	description = "foo"
	other = "bazr"
`

func TestTranslationLoadingE2E(t *testing.T) {
	fs := fstest.MapFS{
		"mock.en.toml": &fstest.MapFile{
			Data:    []byte(mockTranslation),
			Mode:    fs.ModeAppend,
			ModTime: time.Now(),
			Sys:     nil,
		},
	}

	bundle, err := Load(fs, language.English)
	if err != nil {
		t.Error(err)
		return
	}

	tags := bundle.LanguageTags()
	if len(tags) != 1 {
		t.Error("expected bundle to contain exactly one language tag")
	}
	if tags[0] != language.English {
		t.Error("expected bundle to english language tag")
	}
}
