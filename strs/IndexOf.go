package strs

import (
	"gitlab.com/evatix-go/strhelper/internal/pkg/panichelper"
	"gitlab.com/evatix-go/strhelper/strconst"
)

// Returns the index where the string first found, rest don't care
func IndexOf(
	lines *[]string,
	findingString *string,
	startsAtIndex int,
	isCaseSensitive bool,
) int {
	if IsEmpty(lines) || findingString == nil {
		return strconst.InvalidNotFoundCase
	}

	length := len(*lines)

	if startsAtIndex <= strconst.InvalidNotFoundCase || startsAtIndex > length-1 {
		panichelper.StartAtIndexFailed(startsAtIndex, length)
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

	return strconst.InvalidNotFoundCase
}
