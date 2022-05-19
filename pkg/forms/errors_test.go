package forms

import (
	"testing"
)

func TestErrors(t *testing.T) {
	e := errors{localizer: newLocalizer(t)}
	e.Add("foo", "FooError")

	want := "foo"
	got, err := e.Get("foo")
	if err != nil {
		t.Fatal(err)
	}

	if want != got {
		t.Errorf("want %s; got %s on error map", want, got)
	}

	e.Add("foo", "BarError")
	got, err = e.Get("foo")
	if err != nil {
		t.Fatal(err)
	}

	if want != got {
		t.Errorf("want %s; got %s on error map", want, got)
	}
}
