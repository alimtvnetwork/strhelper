package strs

import (
	"strings"

	"gitlab.com/evatix-go/strhelper/constants"
)

func IndexOfForCaseInsensitive(
	lines *[]string,
	findingString *string,
	startsAtIndex int,
) int {
	if IsEmpty(lines) || findingString == nil {
		return constants.InvalidNotFoundCase
	}

	length := len(*lines)

	if startsAtIndex <= constants.InvalidNotFoundCase || startsAtIndex > length-1 {
		message := "startsAtIndex cannot be negative or more than length. startsAtIndex:" + string(startsAtIndex)

		panic(message)
	}

	findingStringToLower := strings.ToLower(*findingString)

	for i := startsAtIndex; i < length; i++ {
		if strings.ToLower((*lines)[i]) == findingStringToLower {
			return i
		}
	}

	return constants.InvalidNotFoundCase
}

// Assumptions are lines, findingString are check already not null or empty
// Kept for internal use only.
func indexOfForCaseInsensitive(lines *[]string, findingString *string, startsAtIndex int) int {
	length := len(*lines)
	findingStringToLower := strings.ToLower(*findingString)

	for i := startsAtIndex; i < length; i++ {
		if strings.ToLower((*lines)[i]) == findingStringToLower {
			return i
		}
	}

	return constants.InvalidNotFoundCase
}
