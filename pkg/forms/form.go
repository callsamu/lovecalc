package forms

import (
	"net/url"
	"unicode"
	"unicode/utf8"

	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type Form struct {
	url.Values
	Errors errors
}

func New(values url.Values, l *i18n.Localizer) *Form {
	return &Form{
		Values: values,
		Errors: errors{
			localizer: l,
			configs:   map[string][]*i18n.LocalizeConfig{},
		},
	}
}

func (f *Form) Required(fields ...string) *Form {
	for _, field := range fields {
		if f.Get(field) == "" {
			f.Errors.Add(field, "FieldIsRequiredError")
		}
	}

	return f
}

func (f *Form) MaxLength(field string, max int) *Form {
	if utf8.RuneCountInString(field) > max {
		f.Errors.Addc(field, "FieldTooLongError", max)
	}

	return f
}

func (f *Form) UnicodeLettersOnly(fields ...string) *Form {
	for _, field := range fields {
		for _, rune := range f.Get(field) {
			if !(unicode.IsLetter(rune) || unicode.IsSpace(rune)) {
				f.Errors.Add(field, "FieldHasInvalidCharsError")
			}
		}
	}

	return f
}

func (f *Form) Valid() bool {
	return f.Errors.len() == 0
}
