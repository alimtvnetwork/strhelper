package strsindex

import (
	"strings"

	"gitlab.com/evatix-go/strhelper/internal/isstrsinternal"
	"gitlab.com/evatix-go/strhelper/internal/panichelper"
	"gitlab.com/evatix-go/strhelper/strconst"
)

func OfCaseInsensitive(
	lines *[]string,
	findingString *string,
	startsAtIndex int,
) int {
	if isstrsinternal.EmptyPtr(lines) || findingString == nil {
		return strconst.InvalidNotFoundCase
	}

	length := len(*lines)

	if startsAtIndex <= strconst.InvalidNotFoundCase || startsAtIndex > length-1 {
		panichelper.StartAtIndexFailed(startsAtIndex, length)
	}

	findingStringToLower := strings.ToLower(*findingString)

	for i := startsAtIndex; i < length; i++ {
		if strings.ToLower((*lines)[i]) == findingStringToLower {
			return i
		}
	}

	return strconst.InvalidNotFoundCase
}
