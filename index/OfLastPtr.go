package index

import (
	"strings"

	"gitlab.com/evatix-go/strhelper/internal/pkg/constants"
)

// Returns the last index of the findingString in s
// it returns the index where the word starts from not the end of index
// startsAt cannot be negative
// If found returns the index from last, if not then returns -1
func OfLastPtr(s, findingString *string, startsAt int, isCaseSensitive bool) int {
	if s == nil || findingString == nil {
		panic(constants.SearchNullPanicMessage)
	}

	if isCaseSensitive && startsAt == 0 {
		return strings.LastIndex(*s, *findingString)
	}

	if isCaseSensitive {
		return OfLastCaseSensitive(s, findingString, startsAt)
	}

	return OfLastCaseInsensitive(s, findingString, startsAt)
}
