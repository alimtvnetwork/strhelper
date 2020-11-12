package strs

import "gitlab.com/evatix-go/strhelper/constants"

// Returns the index where the string first found, rest don't care
func IndexOfPtrStr(
	lines *[]*string,
	findingString *string,
	startsAtIndex int,
	isCaseSensitive bool,
) int {
	if IsEmptyPtrStr(lines) || findingString == nil {
		return constants.InvalidNotFoundCase
	}

	length := len(*lines)

	if startsAtIndex <= constants.InvalidNotFoundCase || startsAtIndex > length-1 {
		startAtIndexFailed(startsAtIndex, length)
	}

	if !isCaseSensitive {
		// insensitive
		return indexOfPtrStrForCaseInsensitiveInternal(
			lines,
			findingString,
			startsAtIndex)
	}

	for ; startsAtIndex < length; startsAtIndex++ {
		if *(*lines)[startsAtIndex] == *findingString {
			return startsAtIndex
		}
	}

	return constants.InvalidNotFoundCase
}
