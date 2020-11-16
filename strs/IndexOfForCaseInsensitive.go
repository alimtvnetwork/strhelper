package strs

import (
	"strings"

	"gitlab.com/evatix-go/strhelper/internal/pkg/panichelper"
	"gitlab.com/evatix-go/strhelper/strconst"
)

func IndexOfForCaseInsensitive(
	lines *[]string,
	findingString *string,
	startsAtIndex int,
) int {
	if IsEmpty(lines) || findingString == nil {
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
