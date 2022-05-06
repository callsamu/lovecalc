package namemixer

import "testing"

func TestMixer(t *testing.T) {
	cases := []struct {
		first  string
		second string
		want   string
	}{
		{"Delmiro", "Valeria", "Delmiria"},
		{"Joao", "Ana", "Joaana"},
	}

	for _, tc := range cases {
		got := MixNames(tc.first, tc.second)
		if tc.want != got {
			t.Errorf("want %s; got %s", tc.want, got)
		}
	}
}
