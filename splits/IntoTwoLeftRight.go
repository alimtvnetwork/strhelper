package splits

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/corestr"
	"gitlab.com/evatix-go/core/coreindexes"
	"gitlab.com/evatix-go/core/errcore"
)

func IntoTwoLeftRight(s, separator string) corestr.LeftRight {
	splits := strings.SplitN(
		s, separator,
		ExpectingLengthOfIntoTwoSplits)

	length := len(splits)
	first := splits[coreindexes.First]

	if length == ExpectingLengthOfIntoTwoSplits {
		return corestr.LeftRight{
			Left:    first,
			Right:   splits[coreindexes.Second],
			IsValid: true,
			Message: "",
		}
	}

	return corestr.LeftRight{
		Left:    first,
		Right:   constants.EmptyString,
		IsValid: false,
		Message: errcore.Expecting(
			"Expecting length",
			ExpectingLengthOfIntoTwoSplits,
			length),
	}
}
