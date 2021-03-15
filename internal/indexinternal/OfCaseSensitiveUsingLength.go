package indexinternal

import (
	"gitlab.com/evatix-go/core/constants"

	"gitlab.com/evatix-go/strhelper/internal/isstrinternal"
)

// returns -1 on non found case
// panics if any is nil
func OfCaseSensitiveUsingLength(
	s, findingString *string,
	startsAt int,
	wordLength, searchLength int,
) int {
	for i := startsAt; i < searchLength; i++ {
		if searchLength-i < wordLength {
			// there is no need to check anymore
			// exceeded word searchLength and not found case
			break
		}

		if isstrinternal.IsStartsWithUsingLength(
			s,
			findingString,
			i,
			wordLength,
			searchLength) {
			return i
		}
	}

	return constants.InvalidNotFoundCase
}
