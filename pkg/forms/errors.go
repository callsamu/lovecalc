package forms

import "github.com/nicksnyder/go-i18n/v2/i18n"

type errors struct {
	localizer *i18n.Localizer
	configs   map[string][]*i18n.LocalizeConfig
}

func (e *errors) Get(field string) (string, error) {
	es := e.configs[field]
	if len(es) == 0 {
		return "", nil
	}

	return e.localizer.Localize(es[0])
}

// Use message with plural
func (e *errors) Addc(field, message string, count int) {
	cfg := &i18n.LocalizeConfig{
		MessageID:   message,
		PluralCount: count,
		TemplateData: map[string]interface{}{
			"count": count,
		},
	}
	e.configs[field] = append(e.configs[field], cfg)
}

func (e *errors) Add(field, message string) {
	cfg := &i18n.LocalizeConfig{MessageID: message}
	e.configs[field] = append(e.configs[field], cfg)
}

func (e *errors) len() int {
	return len(e.configs)
}
