package splits

import (
	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/coreindexes"
)

func IntoTwoFromLastUsingRune(s string, splitRune rune) (left, right string) {
	splits := LastByRunePtr(
		s,
		splitRune,
		ExpectingLengthOfIntoTwoSplits)

	length := len(splits)
	first := splits[coreindexes.First]

	if length == ExpectingLengthOfIntoTwoSplits {
		return splits[coreindexes.Second], first
	}

	return constants.EmptyString, first
}
