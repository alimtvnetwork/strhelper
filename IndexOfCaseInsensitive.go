package strhelper

import (
	"strings"

	"gitlab.com/evatix-go/strhelper/constants"
)

// returns -1 on non found case
// panics if any is nil
func IndexOfCaseInsensitive(s, findingString *string, startAt int) int {
	if s == nil || findingString == nil {
		panic(searchNullPanicMessage)
	}

	length := len(*s)
	wordLength := len(*findingString)

	if wordLength > length {
		return constants.InvalidNotFoundCase
	}

	strLower := strings.ToLower(*s)
	wordLower := strings.ToLower(*findingString)

	for i := startAt; i < length; i++ {
		if length-i < wordLength {
			// there is no need to check anymore
			// exceeded word length and not found case
			break
		}

		if isStartsWithInternal(&strLower, &wordLower, i) {
			return i
		}
	}

	return constants.InvalidNotFoundCase
}
