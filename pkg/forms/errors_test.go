package forms

import "testing"

func TestErrors(t *testing.T) {
	e := errors{}
	e.Add("foo", "foo_error")

	want := "foo_error"
	got := e.Get("foo")

	if want != got {
		t.Errorf("want %s; got %s on error map %v", want, got, e)
	}

	e.Add("foo", "bar_error")
	got = e.Get("foo")

	if want != got {
		t.Errorf("want %s; got %s on error map %v", want, got, e)
	}
}
