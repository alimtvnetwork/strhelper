package strs

import "gitlab.com/evatix-go/strhelper/constants"

// Returns the index where the string first found, rest don't care
func IndexOf(
	lines *[]string,
	findingString *string,
	startsAtIndex int,
	isCaseSensitive bool,
) int {
	if IsEmpty(lines) || findingString == nil {
		return constants.InvalidNotFoundCase
	}

	length := len(*lines)

	if startsAtIndex <= constants.InvalidNotFoundCase || startsAtIndex > length-1 {
		startAtIndexFailed(startsAtIndex)
	}

	if !isCaseSensitive {
		// insensitive
		return indexOfForCaseInsensitiveInternal(lines, findingString, startsAtIndex)
	}

	for i := startsAtIndex; i < length; i++ {
		if (*lines)[i] == *findingString {
			return i
		}
	}

	return constants.InvalidNotFoundCase
}
