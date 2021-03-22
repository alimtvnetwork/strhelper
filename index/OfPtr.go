package index

import (
	"strings"

	"gitlab.com/evatix-go/strhelper/internal/messages"
	"gitlab.com/evatix-go/strhelper/internal/panichelper"
)

// Returns the first index of the findingString in s
//
// Returns Index
//  - If text is found and nothing is invalid like (none is nil)
//
// Returns -1
//  - When not found or invalid case.
//
// Conditions (for panic):
//  - s or search should NOT be nil.
//  - startsAt cannot be negative.
//  - startsAt larger than the content length.
func OfPtr(s, findingString *string, startsAt int, isCaseSensitive bool) int {
	if s == nil || findingString == nil {
		panic(messages.SearchNullPanicMessage)
	}

	length := len(*s)
	if startsAt < 0 || length-1 < startsAt {
		panichelper.StartAtIndexFailed(startsAt, length)
	}

	if isCaseSensitive && startsAt == 0 {
		return strings.Index(*s, *findingString)
	}

	if isCaseSensitive {
		return OfCaseSensitive(s, findingString, startsAt)
	}

	return OfCaseInsensitive(s, findingString, startsAt)
}
