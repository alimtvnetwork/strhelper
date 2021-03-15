package index

import (
	"gitlab.com/evatix-go/core/constants"

	"gitlab.com/evatix-go/strhelper/internal/isstrinternal"
)

// returns -1 on non found case
// panics if any is nil
func OfLastCaseSensitive(s, findingString *string, lastIndexIncreasedBy int) int {
	wholeTextLength := len(*s)
	searchTextLength := len(*findingString)

	if searchTextLength > wholeTextLength {
		return constants.InvalidNotFoundCase
	}

	for newStartIndex := lastIndexIncreasedBy; newStartIndex < wholeTextLength; newStartIndex++ {
		if wholeTextLength-newStartIndex < searchTextLength {
			// there is no need to check anymore
			// exceeded word wholeTextLength and not found case
			break
		}

		if isstrinternal.IsEndsWith(s, findingString, newStartIndex) {
			return wholeTextLength - newStartIndex - searchTextLength
		}
	}

	return constants.InvalidNotFoundCase
}
