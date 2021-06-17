package splits

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coreindexes"
)

func IntoTwoTrimSpacePtr(s, separator *string) (left, right string) {
	splits := strings.SplitN(
		*s, *separator,
		ExpectingLengthOfIntoTwoSplits)

	length := len(splits)
	first := strings.TrimSpace(splits[coreindexes.First])

	if length == ExpectingLengthOfIntoTwoSplits {
		return first, strings.TrimSpace(splits[coreindexes.Second])
	}

	return first, constants.EmptyString
}
