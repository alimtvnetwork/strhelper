package splits

import (
	"strings"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/coreindexes"
)

func IntoTwoTrimSpace(s, separator string) (left, right string) {
	splits := strings.SplitN(
		s, separator,
		ExpectingLengthOfIntoTwoSplits)

	length := len(splits)
	first := strings.TrimSpace(splits[coreindexes.First])

	if length == ExpectingLengthOfIntoTwoSplits {
		return first, strings.TrimSpace(splits[coreindexes.Second])
	}

	return first, constants.EmptyString
}
