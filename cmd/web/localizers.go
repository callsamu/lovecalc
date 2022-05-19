package main

import (
	"fmt"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

type LocaleManager struct {
	localizers map[string]*i18n.Localizer
	matcher    language.Matcher
}

func NewLocaleManager(bundle *i18n.Bundle) *LocaleManager {
	localizers := map[string]*i18n.Localizer{}
	tags := bundle.LanguageTags()

	for _, tag := range tags {
		lang := tag.String()
		local := i18n.NewLocalizer(bundle, lang)
		localizers[lang] = local
	}

	return &LocaleManager{
		localizers: localizers,
		matcher:    language.NewMatcher(tags),
	}
}

func (lm *LocaleManager) localize(key string, td *templateData) (string, error) {
	lang := td.Lang
	local, ok := lm.localizers[lang]
	if !ok {
		return "", fmt.Errorf("locale \"%s\" wasn't found", lang)
	}

	message, err := local.Localize(&i18n.LocalizeConfig{
		MessageID:    key,
		TemplateData: td,
	})

	return message, err
}
