package index

import (
	"gitlab.com/evatix-go/core/constants"

	"gitlab.com/evatix-go/strhelper/internal/isstrinternal"
)

// returns -1 on non found case
// panics if any is nil
func ofLastCaseSensitiveUsingLength(
	s, findingString *string,
	contentLengthDecreasedBy int,
	wholeTextLength, searchLength int,
) int {
	// it will normally go as OfIndex, 0.1.2.3...N
	for newStartIndex := contentLengthDecreasedBy; newStartIndex < wholeTextLength; newStartIndex++ {
		if wholeTextLength-newStartIndex < searchLength {
			// there is no need to check anymore
			// exceeded word wholeTextLength and not found case
			break
		}

		// here having newStartIndex = 1 will compare from last index - newStartIndex
		if isstrinternal.EndsWithUsingLength(
			s, findingString,
			newStartIndex,
			wholeTextLength,
			searchLength) {
			return wholeTextLength - newStartIndex - searchLength
		}
	}

	return constants.InvalidNotFoundCase
}
