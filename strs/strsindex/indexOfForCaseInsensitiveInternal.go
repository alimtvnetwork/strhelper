package strsindex

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
)

// Assumptions are lines, findingString are checked already not null or empty
// Kept for internal use only.
func indexOfForCaseInsensitiveInternal(
	lines []string,
	findingString string,
	startsAtIndex int,
) int {
	length := len(lines)
	for i := startsAtIndex; i < length; i++ {
		if strings.EqualFold(lines[i], findingString) {
			return i
		}
	}

	return constants.InvalidNotFoundCase
}
