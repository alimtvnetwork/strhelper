package strhelper

import (
	"strings"

	"gitlab.com/evatix-go/strhelper/constants"
)

// startsAtLastIndex = 0 meaning starts from last position, giving 2 meaning len(wholeText)-2
// If EmptyString(constants.EmptyString) is given for search and if the startsAt less than the length of the wholeText then it returns true.
// Warning : If any of the searching strings are given nil then it will panic
func IsEndsWithPtr(
	wholeText, endsWithSearch *string,
	startsAtLastIndex int,
	isCaseSensitive bool,
) bool {
	if wholeText == nil || endsWithSearch == nil {
		panic("Either wholeText or endsWithSearch is nil. Please provide valid string at least EmptyString(\"\")")
	}

	searchLength := len(*endsWithSearch)
	wholeTextLength := len(*wholeText)

	if searchLength == constants.Zero {
		return (wholeTextLength == constants.Zero && startsAtLastIndex == constants.Zero) || wholeTextLength-1 >= startsAtLastIndex
	}

	if wholeTextLength == constants.Zero {
		return searchLength == constants.Zero && startsAtLastIndex == constants.Zero
	}

	textLength := wholeTextLength - startsAtLastIndex
	if searchLength > textLength {
		return false
	}

	startingIndex := textLength - searchLength
	substringFromText := (*wholeText)[startingIndex:textLength]

	if isCaseSensitive {
		return substringFromText == *endsWithSearch
	}

	lowerCaseSubstring := strings.ToLower(substringFromText)
	lowerCaseEndSearchText := strings.ToLower(*endsWithSearch)

	// insensitive
	return lowerCaseSubstring == lowerCaseEndSearchText
}
