package main

import (
	"errors"
	"fmt"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var errUnsupportedLocale = errors.New("unsupported locale")

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

func (lm *LocaleManager) GetLocalizer(lang string) (*i18n.Localizer, error) {
	l, ok := lm.localizers[lang]
	if !ok {
		return nil, fmt.Errorf("%w: %s", lang)
	}
	return l, nil
}

func (lm *LocaleManager) Match(lang ...string) language.Tag {
	match, _ := language.MatchStrings(lm.matcher, lang...)
	return match
}
