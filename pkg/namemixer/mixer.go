package namemixer

import "strings"

func MixNames(first, second string) string {
	syllabes := splitSyllabes(first)
	firstHalf := syllabes
	if len(syllabes) > 2 {
		firstHalf = firstHalf[:2]
	}

	syllabes = splitSyllabes(strings.ToLower(second))
	secondHalf, length := syllabes, len(syllabes)
	if length > 2 {
		secondHalf = secondHalf[length-2 : length]
	}

	syllabes = append(firstHalf, secondHalf...)
	return strings.Join(syllabes, "")
}
