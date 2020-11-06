package strhelper

import (
	"strings"

	"gitlab.com/evatix-go/strhelper/constants"
)

// startsAt = 0 meaning starts from last position, giving 2 meaning start index += 2
// If EmptyString(constants.EmptyString) is given for search and if the startsAt less than the length of the wholeText then it returns true.
// Warning : If any of the searching strings are given nil then it will panic
func IsStartsWithPtr(
	wholeText, startsWith *string,
	startsAt int,
	isCaseSensitive bool,
) bool {
	if wholeText == nil || startsWith == nil {
		panic("Either wholeText or startsWith is nil. Please provide valid string at least EmptyString(\"\")")
	}

	searchLength := len(*startsWith)
	wholeTextLength := len(*wholeText)

	if searchLength == constants.Zero {
		return wholeTextLength == constants.Zero && startsAt == constants.Zero || wholeTextLength-1 >= startsAt
	}

	if wholeTextLength == constants.Zero {
		return searchLength == constants.Zero && startsAt == constants.Zero
	}

	textLength := wholeTextLength - startsAt

	if searchLength > textLength {
		return false
	}

	endingLength := startsAt + searchLength
	substringFromText := (*wholeText)[startsAt:endingLength]

	if isCaseSensitive {
		return substringFromText == *startsWith
	}

	lowerCaseSubstring := strings.ToLower(substringFromText)
	lowerCaseSearchText := strings.ToLower(*startsWith)

	// insensitive
	return lowerCaseSubstring == lowerCaseSearchText
}
