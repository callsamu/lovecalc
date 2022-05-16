package main

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

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
		local := localizers[lang]

		message, err := local.Localize(&i18n.LocalizeConfig{
			MessageID:    key,
			TemplateData: td,
		})

		return message, err
	}
}
