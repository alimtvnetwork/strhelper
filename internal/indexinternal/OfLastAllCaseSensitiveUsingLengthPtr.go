package indexinternal

import (
	"gitlab.com/evatix-go/core/constants"

	"gitlab.com/evatix-go/strhelper/internal/defaultcapacity"
	"gitlab.com/evatix-go/strhelper/internal/isstrinternal"
)

func OfLastAllCaseSensitiveUsingLengthPtr(
	s, findingString *string,
	contentLengthDecreasedBy int,
	limits int,
	wholeTextLength, searchLength int,
) *[]int {
	if wholeTextLength-contentLengthDecreasedBy < searchLength {
		// exceeded word wholeTextLength and not found case
		return nil
	}

	defaultCapacity := defaultcapacity.Get(wholeTextLength, limits)
	indexes := make([]int, 0, defaultCapacity)
	hasLimit := limits > constants.InvalidValue
	foundIndex := constants.InvalidNotFoundCase

	// it will normally go as OfIndex, 0.1.2.3...N
	for newStartIndex := contentLengthDecreasedBy; newStartIndex < wholeTextLength; newStartIndex++ {
		if wholeTextLength-newStartIndex < searchLength || hasLimit && limits <= 0 {
			// there is no need to check anymore
			// exceeded word wholeTextLength and not found case
			break
		}

		// here having newStartIndex = 1 will compare from last index - newStartIndex
		if isstrinternal.EndsWithUsingLength(
			s, findingString,
			newStartIndex,
			wholeTextLength,
			searchLength) {
			foundIndex = wholeTextLength - newStartIndex - searchLength
			indexes = append(indexes, foundIndex)
			limits--
		}
	}

	if foundIndex == constants.InvalidNotFoundCase {
		return nil
	}

	return &indexes
}
