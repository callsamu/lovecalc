package namemixer

import (
	"strings"
)

var vowels = map[rune]struct{}{
	'a': {}, 'e': {}, 'i': {}, 'o': {}, 'u': {},
	'A': {}, 'E': {}, 'I': {}, 'O': {}, 'U': {},
}

func isVowel(letter rune) bool {
	_, isVowel := vowels[letter]
	if isVowel {
		return true
	}

	return false
}

func nextVowel(runes []rune, start int) (int, bool) {
	for i := start; i < len(runes); i++ {
		if isVowel(runes[i]) {
			return i, true
		}
	}

	return 0, false
}

func removeDoubleVowels(s string) string {
	if s == "" {
		return s
	}

	input := []rune(s)
	result := []rune{}

	lastIsVowel := false

	for _, r := range input {
		currIsVowel := isVowel(r)

		if !(currIsVowel && lastIsVowel) {
			result = append(result, r)
		}

		lastIsVowel = currIsVowel
	}

	return string(result)
}

func splitSyllabes(word string) []string {
	word = strings.ReplaceAll(word, " ", "")
	syllabes := []string{}
	runes := []rune(word)

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
