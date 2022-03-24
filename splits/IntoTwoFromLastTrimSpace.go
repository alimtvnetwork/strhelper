package splits

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coreindexes"
)

func IntoTwoFromLastTrimSpace(
	s, separator string, isCaseSensitive bool,
) (left, right string) {
	splits := LastByLimit(
		s,
		separator,
		isCaseSensitive,
		ExpectingLengthOfIntoTwoSplits)

	length := len(splits)
	first := strings.TrimSpace(splits[coreindexes.First])

	if length == ExpectingLengthOfIntoTwoSplits {
		return strings.TrimSpace(splits[coreindexes.Second]), first
	}

	return constants.EmptyString, first
}
