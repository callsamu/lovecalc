package main

import "github.com/nicksnyder/go-i18n/v2/i18n"

type LocalizerFunc func(string, *templateData)

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

