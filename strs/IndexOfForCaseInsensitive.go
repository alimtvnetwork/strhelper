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
		startAtIndexFailed(startsAtIndex)
	}

	findingStringToLower := strings.ToLower(*findingString)

	for i := startsAtIndex; i < length; i++ {
		if strings.ToLower((*lines)[i]) == findingStringToLower {
			return i
		}
	}

	return constants.InvalidNotFoundCase
}
