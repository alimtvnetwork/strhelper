package index

import (
	"gitlab.com/evatix-go/core/constants"
)

// Returns the last index of the findingString in s,
// it returns the index where the word starts from not the end of index
// startsAt cannot be negative
// If found returns the index from last, if not then returns -1
// Use Ptr version for performance
func OfLast(s, findingString string, startsAt int, isCaseSensitive bool) int {
	if s == findingString && startsAt == 0 {
		return constants.Zero
	}

	return OfLastPtr(
		s,
		findingString,
		startsAt,
		isCaseSensitive)
}
