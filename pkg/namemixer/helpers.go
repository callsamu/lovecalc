package namemixer

import "strings"

func getFirstName(fullname string) string {
	names := strings.SplitAfterN(fullname, " ", 2)

	if names != nil {
		return names[0]
	}

	return ""
}
