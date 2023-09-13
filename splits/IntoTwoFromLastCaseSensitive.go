package splits

import (
	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/coreindexes"
)

func IntoTwoFromLastCaseSensitive(
	s, separator string,
) (left, right string) {
	splits := LastByLimit(
		s,
		separator,
		true,
		ExpectingLengthOfIntoTwoSplits)

	length := len(splits)
	first := splits[coreindexes.First]

	if length == ExpectingLengthOfIntoTwoSplits {
		return splits[coreindexes.Second], first
	}

	return constants.EmptyString, first
}
