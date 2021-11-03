package stringindex

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"

	"gitlab.com/evatix-go/strhelper/internal/isstrinternal"
)

// returns -1 on non found case
// panics if any is nil
func OfCaseInsensitive(
	s, findingString string,
	startsAt int,
) int {
	wholeTextLength := len(s)
	searchingLength := len(findingString)

	if searchingLength > wholeTextLength {
		return constants.InvalidNotFoundCase
	}

	strLower := strings.ToLower(s)
	wordLower := strings.ToLower(findingString)

	for i := startsAt; i < wholeTextLength; i++ {
		if wholeTextLength-i < searchingLength {
			// there is no need to check anymore
			// exceeded word wholeTextLength and not found case
			break
		}

		if isstrinternal.StartsWithUsingLength(
			strLower,
			wordLower,
			i,
			wholeTextLength,
			searchingLength) {
			return i
		}
	}

	return constants.InvalidNotFoundCase
}
