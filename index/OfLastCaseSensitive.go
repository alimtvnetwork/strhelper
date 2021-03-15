package index

import (
	"gitlab.com/evatix-go/core/constants"

	"gitlab.com/evatix-go/strhelper/internal/isstrinternal"
)

// returns -1 on non found case
// panics if any is nil
func OfLastCaseSensitive(s, findingString *string, lastIndexIncreasedBy int) int {
	length := len(*s)
	wordLength := len(*findingString)

	if wordLength > length {
		return constants.InvalidNotFoundCase
	}

	textLength := length - lastIndexIncreasedBy

	for newStartIndex := lastIndexIncreasedBy; newStartIndex < textLength; newStartIndex++ {
		if textLength-newStartIndex < wordLength {
			// there is no need to check anymore
			// exceeded word length and not found case
			break
		}

		if isstrinternal.IsEndsWithInternal(s, findingString, newStartIndex) {
			return length - newStartIndex - wordLength
		}
	}

	return constants.InvalidNotFoundCase
}
