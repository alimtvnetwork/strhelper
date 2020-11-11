package strhelper

import (
	"strings"

	"gitlab.com/evatix-go/strhelper/constants"
)

// returns -1 on non found case
// panics if any is nil
func LastIndexOfCaseInsensitive(s, findingString *string, startAt int) int {
	if s == nil || findingString == nil {
		panic(searchNullPanicMessage)
	}

	length := len(*s)
	wordLength := len(*findingString)

	if wordLength > length {
		return constants.InvalidNotFoundCase
	}

	wholeTextLower := strings.ToLower(*s)
	wordLower := strings.ToLower(*findingString)

	textLength := length - startAt

	for newStartIndex := startAt; newStartIndex < textLength; newStartIndex++ {
		if textLength-newStartIndex < wordLength {
			// there is no need to check anymore
			// exceeded word length and not found case
			break
		}

		if isEndsWithInternal(&wholeTextLower, &wordLower, newStartIndex) {
			return length - newStartIndex - wordLength
		}
	}

	return constants.InvalidNotFoundCase
}
