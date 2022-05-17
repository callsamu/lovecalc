package main

import (
	"errors"

	"github.com/nicksnyder/go-i18n/v2/i18n"
)

var errLocalizerNotFound = errors.New("localizers: localizer not found")

type LocalizerFunc func(string, *templateData) (string, error)

func newLocalizers(bundle *i18n.Bundle) map[string]*i18n.Localizer {
	localizers := map[string]*i18n.Localizer{}
	tags := bundle.LanguageTags()

	for _, tag := range tags {
		lang := tag.String()
		local := i18n.NewLocalizer(bundle, lang)
		localizers[lang] = local
	}

	return localizers
}

func newLocalizerFunc(localizers map[string]*i18n.Localizer) LocalizerFunc {
	return func(key string, td *templateData) (string, error) {
		lang := td.Lang
		local, ok := localizers[lang]
		if !ok {
			return "", errLocalizerNotFound
		}

		message, err := local.Localize(&i18n.LocalizeConfig{
			MessageID:    key,
			TemplateData: td,
		})

		return message, err
	}
}
