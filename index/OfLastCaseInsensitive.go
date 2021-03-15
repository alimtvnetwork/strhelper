package index

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"

	"gitlab.com/evatix-go/strhelper/internal/consts"
	"gitlab.com/evatix-go/strhelper/internal/isstrinternal"
)

// returns -1 on non found case
// panics if any is nil
func OfLastCaseInsensitive(s, findingString *string, startsAt int) int {
	if s == nil || findingString == nil {
		panic(consts.SearchNullPanicMessage)
	}

	wholeTextLength := len(*s)
	searchTextLength := len(*findingString)

	if searchTextLength > wholeTextLength {
		return constants.InvalidNotFoundCase
	}

	if s == findingString && startsAt == 0 {
		return constants.Zero
	}

	if *s == *findingString && startsAt == 0 {
		return constants.Zero
	}

	wholeTextLower := strings.ToLower(*s)
	wordLower := strings.ToLower(*findingString)

	for newStartIndex := startsAt; newStartIndex < wholeTextLength; newStartIndex++ {
		if textLength-newStartIndex < searchTextLength {
			// there is no need to check anymore
			// exceeded word wholeTextLength and not found case
			break
		}

		if isstrinternal.IsEndsWith(&wholeTextLower, &wordLower, newStartIndex) {
			return wholeTextLength - newStartIndex - searchTextLength
		}
	}

	return constants.InvalidNotFoundCase
}
