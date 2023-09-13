package strsindex

import (
	"strings"

	"gitlab.com/auk-go/core/constants"

	"gitlab.com/auk-go/strhelper/internal/isstrsinternal"
	"gitlab.com/auk-go/strhelper/internal/panichelper"
)

func OfCaseInsensitive(
	lines *[]string,
	findingString string,
	startsAtIndex int,
) int {
	if isstrsinternal.EmptyPtr(lines) {
		return constants.InvalidNotFoundCase
	}

	length := len(*lines)

	if startsAtIndex <= constants.InvalidNotFoundCase || startsAtIndex > length-1 {
		panichelper.StartAtIndexFailed(startsAtIndex, length)
	}

	linesNonPtr := *lines

	for i := startsAtIndex; i < length; i++ {
		if strings.EqualFold(linesNonPtr[i], findingString) {
			return i
		}
	}

	return constants.InvalidNotFoundCase
}
