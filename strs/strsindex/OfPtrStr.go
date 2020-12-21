package strsindex

import (
	"gitlab.com/evatix-go/strhelper/internal/isstrsinternal"
	"gitlab.com/evatix-go/strhelper/internal/panichelper"
	"gitlab.com/evatix-go/strhelper/strconst"
)

// Returns the index where the string first found, rest don't care
func OfPtrStr(
	lines *[]*string,
	findingString *string,
	startsAtIndex int,
	isCaseSensitive bool,
) int {
	if isstrsinternal.EmptyPtrStr(lines) || findingString == nil {
		return strconst.InvalidNotFoundCase
	}

	length := len(*lines)

	if startsAtIndex <= strconst.InvalidNotFoundCase || startsAtIndex > length-1 {
		panichelper.StartAtIndexFailed(startsAtIndex, length)
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

	return strconst.InvalidNotFoundCase
}
