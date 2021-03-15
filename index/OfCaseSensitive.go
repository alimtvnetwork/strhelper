package index

import (
	"gitlab.com/evatix-go/core/constants"

	"gitlab.com/evatix-go/strhelper/internal/consts"
	"gitlab.com/evatix-go/strhelper/internal/isstrinternal"
)

// returns -1 on non found case
// panics if any is nil
func OfCaseSensitive(s, findingString *string, startsAt int) int {
	if s == nil || findingString == nil {
		panic(consts.SearchNullPanicMessage)
	}

	wholeTextLength := len(*s)
	searchingLength := len(*findingString)

	if searchingLength > wholeTextLength {
		return constants.InvalidNotFoundCase
	}

	for i := startsAt; i < wholeTextLength; i++ {
		if wholeTextLength-i < searchingLength {
			// there is no need to check anymore
			// exceeded word wholeTextLength and not found case
			break
		}

		if isstrinternal.IsStartsWithUsingLength(
			s,
			findingString,
			i,
			wholeTextLength,
			searchingLength) {
			return i
		}
	}

	return constants.InvalidNotFoundCase
}
