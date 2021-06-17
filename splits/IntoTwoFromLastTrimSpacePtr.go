package splits

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coreindexes"
)

func IntoTwoFromLastTrimSpacePtr(s, separator *string, isCaseSensitive bool) (left, right string) {
	splits := LastByLimitPtr(
		s,
		separator,
		isCaseSensitive,
		ExpectingLengthOfIntoTwoSplits)

	length := len(*splits)
	first := strings.TrimSpace((*splits)[coreindexes.First])

	if length == ExpectingLengthOfIntoTwoSplits {
		return strings.TrimSpace((*splits)[coreindexes.Second]), first
	}

	return constants.EmptyString, first
}
