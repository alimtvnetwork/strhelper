package index

import (
	"gitlab.com/evatix-go/core/constants"

	"gitlab.com/evatix-go/strhelper/internal/consts"
	"gitlab.com/evatix-go/strhelper/internal/isstrinternal"
)

// returns -1 on non found case
// panics if any is nil
func OfCaseSensitive(s, findingString *string, startAt int) int {
	if s == nil || findingString == nil {
		panic(consts.SearchNullPanicMessage)
	}

	searchLength := len(*s)
	wordLength := len(*findingString)

	if wordLength > searchLength {
		return constants.InvalidNotFoundCase
	}

	for i := startAt; i < searchLength; i++ {
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
