package strsindex

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
)

// Assumptions are lines, findingString are check already not null or empty
// Kept for internal use only.
func indexOfPtrStrForCaseInsensitiveInternal(lines *[]*string, findingString *string, startsAtIndex int) int {
	length := len(*lines)
	findingStringToLower := strings.ToLower(*findingString)

	for ; startsAtIndex < length; startsAtIndex++ {
		if strings.ToLower(*(*lines)[startsAtIndex]) == findingStringToLower {
			return startsAtIndex
		}
	}

	return constants.InvalidNotFoundCase
}
