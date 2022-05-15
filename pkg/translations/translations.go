package translations

import (
	"embed"
	"io/fs"

	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

//go:embed locales
var LocalesFS embed.FS

func Load(lfs fs.FS, defaultLang language.Tag) (*i18n.Bundle, error) {
	bundle := i18n.NewBundle(defaultLang)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)

	err := fs.WalkDir(lfs, ".", func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}

		mf, err := bundle.LoadMessageFileFS(lfs, d.Name())
		if err != nil {
			return err
		}

		return bundle.AddMessages(mf.Tag, mf.Messages...)
	})
	if err != nil {
		return nil, err
	}

	return bundle, nil
}
