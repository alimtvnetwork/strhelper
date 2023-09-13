package splits

import (
	"strings"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/coreindexes"
)

func IntoTwo(s, separator string) (left, right string) {
	splits := strings.SplitN(
		s, separator,
		ExpectingLengthOfIntoTwoSplits)

	length := len(splits)
	first := splits[coreindexes.First]

	if length == ExpectingLengthOfIntoTwoSplits {
		return first, splits[coreindexes.Second]
	}

	return first, constants.EmptyString
}
