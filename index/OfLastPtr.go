package index

import (
	"strings"

	"gitlab.com/evatix-go/strhelper/internal/consts"
)

// Returns the last index of the searchTerm in s
// it returns the index where the word starts from not the end of index
// lastIndexIncreasedBy cannot be negative
// If found returns the index from last, if not then returns -1
func OfLastPtr(s, searchTerm *string, lastIndexIncreasedBy int, isCaseSensitive bool) int {
	if s == nil || searchTerm == nil {
		panic(consts.SearchNullPanicMessage)
	}

	if isCaseSensitive && lastIndexIncreasedBy == 0 {
		return strings.LastIndex(*s, *searchTerm)
	}

	if isCaseSensitive {
		return OfLastCaseSensitive(s, searchTerm, lastIndexIncreasedBy)
	}

	return OfLastCaseInsensitive(s, searchTerm, lastIndexIncreasedBy)
}
