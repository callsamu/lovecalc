package namemixer

import "testing"

func TestVowelFinder(t *testing.T) {
	cases := []struct {
		name     string
		word     string
		start    int
		wantBool bool
		wantPos  int
	}{
		{
			"Find second vowel in Garcia",
			"Garcia",
			3,
			true,
			4,
		},
		{
			"Give false if there's no vowel rest",
			"Rapppz",
			3,
			false,
			0,
		},
	}

	for _, tc := range cases {
		gotPos, gotBool := nextVowel([]rune(tc.word), tc.start)

		if gotBool != tc.wantBool {
			t.Errorf("want %v; got %v", tc.wantBool, gotPos)
		}

		if gotPos != tc.wantPos {
			t.Errorf("want %v; got %v", tc.wantPos, gotPos)
		}
	}
}

func TestDoubleVowelRemover(t *testing.T) {
	cases := []struct {
		name string
		word string
		want string
	}{
		{"Double vowel", "Daal", "Dal"},
		{"Multiple sucessible vowels", "Maaal", "Mal"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := removeDoubleVowels(tc.word)
			if tc.want != got {
				t.Errorf("want %s; got %s", tc.want, got)
			}
		})
	}
}
