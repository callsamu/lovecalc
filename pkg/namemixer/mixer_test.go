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
