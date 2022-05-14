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

func (f *Form) Required(fields ...string) *Form {
	for _, field := range fields {
		if f.Get(field) == "" {
			f.errors.Add(field, "field is required")
		}
	}

	return f
}

func (f *Form) MaxLength(field string, max int) *Form {
	if utf8.RuneCountInString(field) > max {
		f.errors.Add(field, fmt.Sprintf("field must not be longer than %d characters", max))
	}

	return f
}

func (f *Form) UnicodeLettersOnly(fields ...string) *Form {
	for _, field := range fields {
		for _, rune := range []rune(f.Get(field)) {
			if !(unicode.IsLetter(rune) || unicode.IsSpace(rune)) {
				f.errors.Add(field, "field contains invalid characters")
			}
		}
	}

	return f
}

func (f *Form) Valid() bool {
	return len(f.errors) == 0
}
