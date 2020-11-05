package strhelper

import (
	"strings"

	"gitlab.com/evatix-go/strhelper/constants"
)

// startsAtLastIndex = 0 meaning starts from last position, giving 2 meaning len(wholeText)-2
// If EmptyString(constants.EmptyString) is given for search and if the startsAt less than the length of the wholeText then it returns true.
func IsEndsWith(
	wholeText, endsWithSearch string,
	startsAtLastIndex int,
	isCaseSensitive bool,
) bool {
	if IsEmpty(endsWithSearch) {
		return wholeText == constants.EmptyString && startsAtLastIndex == 0 || len(wholeText)-1 >= startsAtLastIndex
	}

	if IsEmpty(wholeText) {
		return endsWithSearch == constants.EmptyString && startsAtLastIndex == 0
	}

	textLength := len(wholeText) - startsAtLastIndex
	searchLength := len(endsWithSearch)
	if searchLength > textLength {
		return false
	}

	startingIndex := textLength - searchLength
	substringFromText := wholeText[startingIndex:textLength]

	if isCaseSensitive {
		return substringFromText == endsWithSearch
	}

	lowerCaseSubstring := strings.ToLower(substringFromText)
	lowerCaseEndSearchText := strings.ToLower(endsWithSearch)

	// insensitive
	return lowerCaseSubstring == lowerCaseEndSearchText
}
