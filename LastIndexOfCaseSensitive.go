package strhelper

import "gitlab.com/evatix-go/strhelper/constants"

// returns -1 on non found case
// panics if any is nil
func LastIndexOfCaseSensitive(s, findingString *string, startAt int) int {
	length := len(*s)
	wordLength := len(*findingString)

	if wordLength > length {
		return constants.InvalidNotFoundCase
	}

	textLength := length - startAt

	for newStartIndex := startAt; newStartIndex < textLength; newStartIndex++ {
		if textLength-newStartIndex < wordLength {
			// there is no need to check anymore
			// exceeded word length and not found case
			break
		}

		if IsEndsWithPtr(s, findingString, newStartIndex, true) {
			return length - newStartIndex - wordLength
		}
	}

	return constants.InvalidNotFoundCase
}
