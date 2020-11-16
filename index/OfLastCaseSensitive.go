package index

import (
	"gitlab.com/evatix-go/strhelper/internal/isinternal"
	"gitlab.com/evatix-go/strhelper/strconst"
)

// returns -1 on non found case
// panics if any is nil
func OfLastCaseSensitive(s, findingString *string, startAt int) int {
	length := len(*s)
	wordLength := len(*findingString)

	if wordLength > length {
		return strconst.InvalidNotFoundCase
	}

	textLength := length - startAt

	for newStartIndex := startAt; newStartIndex < textLength; newStartIndex++ {
		if textLength-newStartIndex < wordLength {
			// there is no need to check anymore
			// exceeded word length and not found case
			break
		}

		if isinternal.IsEndsWithInternal(s, findingString, newStartIndex) {
			return length - newStartIndex - wordLength
		}
	}

	return strconst.InvalidNotFoundCase
}
