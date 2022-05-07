package namemixer

import "strings"

var vowels = map[rune]struct{}{
	'a': {}, 'e': {}, 'i': {}, 'o': {}, 'u': {},
	'A': {}, 'E': {}, 'I': {}, 'O': {}, 'U': {},
}

func nextVowel(runes []rune, start int) (int, bool) {
	for i := start; i < len(runes); i++ {
		_, isVowel := vowels[runes[i]]
		if isVowel {
			return i, true
		}
	}

	return 0, false
}

func splitSyllabes(word string) []string {
	syllabes := []string{}
	runes := []rune(strings.ReplaceAll(word, " ", ""))

	start := 0
	end := 0

	_, hasVowels := nextVowel(runes, start)
	if !hasVowels {
		return []string{word}
	}

	for {
		pos, vowelsLeft := nextVowel(runes, start)
		if !vowelsLeft {
			return syllabes
		}

		peek, hasNextVowel := nextVowel(runes, pos+1)
		if peek > pos+2 {
			end = pos + 2
		} else {
			if !hasNextVowel {
				end = len(runes)
			} else {
				end = pos + 1
			}
		}

		syllabe := word[start:end]
		syllabes = append(syllabes, syllabe)

		start = end
	}
}
