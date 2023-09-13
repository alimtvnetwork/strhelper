package strsindex

import (
	"strings"

	"gitlab.com/auk-go/core/constants"

	"gitlab.com/auk-go/strhelper/internal/isstrsinternal"
	"gitlab.com/auk-go/strhelper/internal/panichelper"
)

// OfLastPtr returns the last index where the string first found, doesn't care about the rest of the items once found.
func OfLastPtr(
	lines *[]string,
	searchTerm string,
	contentLengthDecreasedBy int,
	isCaseSensitive bool,
) int {
	if isstrsinternal.EmptyPtr(lines) {
		return constants.InvalidNotFoundCase
	}

	length := len(*lines)

	if contentLengthDecreasedBy <= constants.InvalidNotFoundCase || contentLengthDecreasedBy > length-1 {
		panichelper.LastIndexIncreasedByFailed(contentLengthDecreasedBy, length)
	}

	linesNonPtr := *lines

	if !isCaseSensitive {
		// insensitive
		index := length - contentLengthDecreasedBy

		for ; index >= 0; index-- {
			if strings.EqualFold(linesNonPtr[index], searchTerm) {
				return index
			}
		}

		return constants.InvalidNotFoundCase
	}

	index := length - 1 - contentLengthDecreasedBy

	for ; index >= 0; index-- {
		if linesNonPtr[index] == searchTerm {
			return index
		}
	}

	return constants.InvalidNotFoundCase
}
