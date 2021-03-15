package index

import (
	"gitlab.com/evatix-go/core/constants"

	"gitlab.com/evatix-go/strhelper/internal/isstrinternal"
)

// returns -1 on non found case
// panics if any is nil
func ofLastCaseSensitiveUsingLength(
	s, findingString *string,
	contentLengthDecreasedBy int,
	wholeTextLength, searchLength int,
) int {
	textLength :=
		wholeTextLength - contentLengthDecreasedBy

	for newStartIndex := contentLengthDecreasedBy; newStartIndex < textLength; newStartIndex++ {
		if textLength-newStartIndex < searchLength {
			// there is no need to check anymore
			// exceeded word wholeTextLength and not found case
			break
		}

		if isstrinternal.IsEndsWith(s, findingString, newStartIndex) {
			return wholeTextLength - newStartIndex - searchLength
		}
	}

	return constants.InvalidNotFoundCase
}
