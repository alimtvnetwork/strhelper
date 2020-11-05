package strhelper

import "gitlab.com/evatix-go/strhelper/constants"

// returns -1 on non found case
// panics if any is nil
func IndexOfCaseSensitive(s, findingString *string, startAt int) int {
	length := len(*s)
	wordLength := len(*findingString)

	if wordLength > length {
		return constants.InvalidNotFoundCase
	}

	for i := startAt; i < length; i++ {
		if length-i < wordLength {
			// there is no need to check anymore
			// exceeded word length and not found case
			break
		}

		if IsStartsWithPtr(s, findingString, startAt, false) {
			return i
		}
	}

	return constants.InvalidNotFoundCase
}
