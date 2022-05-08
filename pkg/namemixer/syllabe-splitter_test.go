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

func TestSyllabeSplitter(t *testing.T) {
	cases := []struct {
		name string
		word string
		want []string
	}{
		{
			"Simple example",
			"Dale",
			[]string{"Da", "le"},
		},
		{
			"More complicated example",
			"Ana",
			[]string{"A", "na"},
		},
		{
			"Even more complicated example",
			"Finland",
			[]string{"Fin", "land"},
		},
		{
			"Brazilian name",
			"Bruno",
			[]string{"Bru", "no"},
		},
		{
			"No vowels",
			"kkkkkk",
			[]string{"kkkkkk"},
		},
		{
			"Double vowels",
			"Hiato",
			[]string{"Hi", "a", "to"},
		},
		{
			"Handle whitespace",
			"Demi Lovato",
			[]string{"De", "mi", "Lo", "va", "to"},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := splitSyllabes(tc.word)

			if len(got) != len(tc.want) {
				t.Errorf("want %v; got %v", tc.want, got)
			}

			for i := range got {
				if tc.want[i] != got[i] {
					t.Errorf("want %v; got %v", tc.want, got)
					return
				}
			}
		})
	}
}
