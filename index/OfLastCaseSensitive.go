package index

import (
	"gitlab.com/evatix-go/core/constants"

	"gitlab.com/evatix-go/strhelper/internal/isstrinternal"
)

// returns -1 on non found case
// panics if any is nil
func OfLastCaseSensitive(s, findingString *string, contentLengthDecreasedBy int) int {
	wholeTextLength := len(*s)
	searchTextLength := len(*findingString)

	if searchTextLength > wholeTextLength {
		return constants.InvalidNotFoundCase
	}

	if s == findingString && contentLengthDecreasedBy == 0 {
		return constants.Zero
	}

	if *s == *findingString && contentLengthDecreasedBy == 0 {
		return constants.Zero
	}

	// it will normally go as OfIndex, 0.1.2.3...N
	for newStartIndex := contentLengthDecreasedBy; newStartIndex < wholeTextLength; newStartIndex++ {
		if wholeTextLength-newStartIndex < searchTextLength {
			// there is no need to check anymore
			// exceeded word wholeTextLength and not found case
			break
		}

		// here having newStartIndex = 1 will compare from last index - newStartIndex
		if isstrinternal.EndsWithUsingLength(
			s, findingString,
			newStartIndex,
			wholeTextLength,
			searchTextLength) {
			return wholeTextLength - newStartIndex - searchTextLength
		}
	}

	return constants.InvalidNotFoundCase
}
