package strsindex

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
)

// Assumptions are lines, findingString are check already not null or empty
// Kept for internal use only.
func indexOfPtrStrForCaseInsensitiveInternal(
	lines *[]*string,
	findingString string,
	startsAtIndex int,
) int {
	linesNonPtr := *lines
	length := len(linesNonPtr)

	for ; startsAtIndex < length; startsAtIndex++ {
		if strings.EqualFold(*linesNonPtr[startsAtIndex], findingString) {
			return startsAtIndex
		}
	}

	return constants.InvalidNotFoundCase
}
