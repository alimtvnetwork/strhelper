package splits

import (
	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/coreindexes"
)

func IntoTwoFromLast(s, separator string, isCaseSensitive bool) (left, right string) {
	splits := LastByLimit(
		s,
		separator,
		isCaseSensitive,
		ExpectingLengthOfIntoTwoSplits)

	length := len(splits)
	first := splits[coreindexes.First]

	if length == ExpectingLengthOfIntoTwoSplits {
		return splits[coreindexes.Second], first
	}

	return constants.EmptyString, first
}
