package strs

import (
	"strings"

	"gitlab.com/evatix-go/strhelper/strconst"
)

// Assumptions are lines, findingString are check already not null or empty
// Kept for internal use only.
func indexOfForCaseInsensitiveInternal(lines *[]string, findingString *string, startsAtIndex int) int {
	length := len(*lines)
	findingStringToLower := strings.ToLower(*findingString)

	for i := startsAtIndex; i < length; i++ {
		if strings.ToLower((*lines)[i]) == findingStringToLower {
			return i
		}
	}

	return strconst.InvalidNotFoundCase
}
