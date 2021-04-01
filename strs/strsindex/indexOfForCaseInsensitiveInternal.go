package strsindex

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
)

// Assumptions are lines, findingString are check already not null or empty
// Kept for internal use only.
func indexOfForCaseInsensitiveInternal(lines *[]string, findingString *string, startsAtIndex int) int {
	length := len(*lines)
	findingStringCopy := *findingString

	for i := startsAtIndex; i < length; i++ {
		if strings.EqualFold((*lines)[i], findingStringCopy) {
			return i
		}
	}

	return constants.InvalidNotFoundCase
}
