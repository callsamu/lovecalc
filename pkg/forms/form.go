package forms

import (
	"fmt"
	"net/url"
	"unicode"
	"unicode/utf8"
)

type Form struct {
	url.Values
	errors errors
}

func New(values url.Values) *Form {
	return &Form{
		Values: values,
		errors: errors{},
	}
}

func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		if f.Get(field) == "" {
			f.errors.Add(field, "field is required")
		}
	}
}

func (f *Form) MaxLength(field string, max int) {
	if utf8.RuneCountInString(field) > max {
		f.errors.Add(field, fmt.Sprintf("field must not be longer than %d characters", max))
	}
}

func (f *Form) UnicodeLettersOnly(field string) {
	for _, rune := range []rune(f.Get(field)) {
		if !(unicode.IsLetter(rune) || unicode.IsSpace(rune)) {
			f.errors.Add(field, "field contains invalid characters")
		}
	}
}

func (f *Form) Valid() bool {
	return len(f.errors) == 0
}
