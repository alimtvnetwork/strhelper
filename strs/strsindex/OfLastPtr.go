package strsindex

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"

	"gitlab.com/evatix-go/strhelper/internal/isstrsinternal"
	"gitlab.com/evatix-go/strhelper/internal/panichelper"
)

// OfLastPtr returns the last index where the string first found, doesn't care about the rest of the items once found.
func OfLastPtr(
	lines *[]string,
	searchTerm *string,
	contentLengthDecreasedBy int,
	isCaseSensitive bool,
) int {
	if isstrsinternal.EmptyPtr(lines) || searchTerm == nil {
		return constants.InvalidNotFoundCase
	}

	length := len(*lines)

	if contentLengthDecreasedBy <= constants.InvalidNotFoundCase || contentLengthDecreasedBy > length-1 {
		panichelper.LastIndexIncreasedByFailed(contentLengthDecreasedBy, length)
	}

	if !isCaseSensitive {
		// insensitive
		toLowerSearchTerm := strings.ToLower(*searchTerm)
		index := length - contentLengthDecreasedBy

		for ; index >= 0; index-- {
			if strings.ToLower((*lines)[index]) == toLowerSearchTerm {
				return index
			}
		}
	}

	index := length - 1 - contentLengthDecreasedBy

	for ; index >= 0; index-- {
		if (*lines)[index] == *searchTerm {
			return index
		}
	}

	return constants.InvalidNotFoundCase
}
