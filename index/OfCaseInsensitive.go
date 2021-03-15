package index

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"

	"gitlab.com/evatix-go/strhelper/internal/consts"
	"gitlab.com/evatix-go/strhelper/internal/isstrinternal"
)

// returns -1 on non found case
// panics if any is nil
func OfCaseInsensitive(s, findingString *string, startsAt int) int {
	if s == nil || findingString == nil {
		panic(consts.SearchNullPanicMessage)
	}

	length := len(*s)
	wordLength := len(*findingString)

	if wordLength > length {
		return constants.InvalidNotFoundCase
	}

	strLower := strings.ToLower(*s)
	wordLower := strings.ToLower(*findingString)

	for i := startsAt; i < length; i++ {
		if length-i < wordLength {
			// there is no need to check anymore
			// exceeded word length and not found case
			break
		}

		if isstrinternal.IsStartsWith(&strLower, &wordLower, i) {
			return i
		}
	}

	return constants.InvalidNotFoundCase
}
