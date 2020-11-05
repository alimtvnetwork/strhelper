package strhelper

import (
	"strings"
)

// startsAt cannot be negative
func IndexOf(s, findingString *string, startsAt int, isCaseSensitive bool) int {
	if isCaseSensitive && startsAt <= 0 {
		return strings.Index(*s, *findingString)
	}

	if isCaseSensitive {
		return IndexOfCaseSensitive(s, findingString, startsAt)
	}

	return IndexOfCaseInsensitive(s, findingString, startsAt)
}
