package strsindex

import (
	"gitlab.com/evatix-go/core/constants"

	"gitlab.com/evatix-go/strhelper/internal/isstrsinternal"
	"gitlab.com/evatix-go/strhelper/internal/panichelper"
)

// Of Returns the index where the string first found, rest don't care
func Of(
	lines []string,
	findingString string,
	startsAtIndex int,
	isCaseSensitive bool,
) int {
	if isstrsinternal.Empty(lines) {
		return constants.InvalidNotFoundCase
	}

	length := len(lines)

	if startsAtIndex <= constants.InvalidNotFoundCase || startsAtIndex > length-1 {
		panichelper.StartAtIndexFailed(startsAtIndex, length)
	}

	if !isCaseSensitive {
		// insensitive
		return indexOfForCaseInsensitiveInternal(
			lines,
			findingString,
			startsAtIndex)
	}

	for i := startsAtIndex; i < length; i++ {
		if lines[i] == findingString {
			return i
		}
	}

	return constants.InvalidNotFoundCase
}
