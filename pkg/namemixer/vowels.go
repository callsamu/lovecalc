package namemixer

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
