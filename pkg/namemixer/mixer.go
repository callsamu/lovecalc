package namemixer

import (
	"strings"
)

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

func MixNames(first, second string) string {
	syllabes := splitSyllabes(first)
	firstHalf := syllabes
	if len(syllabes) > 2 {
		firstHalf = firstHalf[:2]
	}

	syllabes = splitSyllabes(strings.ToLower(getFirstName(second)))
	secondHalf, length := syllabes, len(syllabes)
	if length > 2 {
		secondHalf = secondHalf[length-2 : length]
	}

	syllabes = append(firstHalf, secondHalf...)
	return strings.Join(syllabes, "")
}
