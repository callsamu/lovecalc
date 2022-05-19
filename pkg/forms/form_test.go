package forms

import (
	"net/url"
	"testing"
)

func TestRequire(t *testing.T) {
	l := newLocalizer(t)

	t.Run("do not supply field", func(t *testing.T) {
		form := New(url.Values{
			"foo": []string{""},
		}, l)

		form.Required("foo")

		if form.Valid() {
			t.Error("incorrect validation: required field not supplied")
		}
	})

	t.Run("supply field", func(t *testing.T) {
		form := New(url.Values{
			"foo": []string{"foo"},
		}, l)

		form.Required("foo")

		if !form.Valid() {
			t.Error("incorrect validation: requirements were fullfiled, but form remains invalid")
		}
	})

	t.Run("supply one field and leave other empty", func(t *testing.T) {
		form := New(url.Values{
			"foo": []string{"foo"},
			"bar": []string{""},
		}, l)

		form.Required("foo", "bar")

		if form.Valid() {
			t.Error("incorrect validation: required field not supplied")
		}
	})
}

func TestMaxLength(t *testing.T) {
	l := newLocalizer(t)

	form := New(url.Values{
		"foo": []string{"foo"},
		"bar": []string{"bar"},
	}, l)

	form.MaxLength("foo", 3)
	if !form.Valid() {
		t.Error("incorrect validation: expected form to be valid")
	}

	form.MaxLength("bar", 1)
	if form.Valid() {
		t.Error("incorrect validation: expected form to be invalid")
	}
}

func TestUnicodeLettersOnly(t *testing.T) {
	l := newLocalizer(t)

	cases := []struct {
		name  string
		field string
		input string
		valid bool
	}{
		{"simple name", "s", "sam", true},
		{"name with spaces", "f", "fox hale", true},
		{"invalid name", "i", "ahmed##", false},
		{"chinese string", "c", "的是了我", true},
		{"string with umlaut", "e", "für elise", true},
	}

	for _, ts := range cases {
		t.Run(ts.name, func(t *testing.T) {
			form := New(url.Values{
				ts.field: []string{ts.input},
			}, l)

			form.UnicodeLettersOnly(ts.field)
			if form.Valid() != ts.valid {
				t.Error("")
			}
		})
	}
}
