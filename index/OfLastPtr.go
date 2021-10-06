package index

import (
	"strings"
)

// Returns the last index of the searchTerm in s
// it returns the index where the word starts from not the end of index
// contentLengthDecreasedBy cannot be negative
// If found returns the index from last, if not then returns -1
func OfLastPtr(
	s, searchTerm string,
	contentLengthDecreasedBy int,
	isCaseSensitive bool,
) int {
	if isCaseSensitive && contentLengthDecreasedBy == 0 {
		return strings.LastIndex(s, searchTerm)
	}

	if isCaseSensitive {
		return OfLastCaseSensitive(s, searchTerm, contentLengthDecreasedBy)
	}

	return OfLastCaseInsensitive(s, searchTerm, contentLengthDecreasedBy)
}
