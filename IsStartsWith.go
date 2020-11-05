package strhelper

import (
	"strings"

	"gitlab.com/evatix-go/strhelper/constants"
)

// startsAt = 0 meaning starts from last position, giving 2 meaning start index += 2
// If EmptyString(constants.EmptyString) is given for search and if the startsAt less than the length of the wholeText then it returns true.
func IsStartsWith(
	wholeText, startsWith string,
	startsAt int,
	isCaseSensitive bool,
) bool {
	if IsEmpty(startsWith) {
		return wholeText == constants.EmptyString && startsAt == 0 || len(wholeText)-1 >= startsAt
	}

	if IsEmpty(wholeText) {
		return startsWith == constants.EmptyString && startsAt == 0
	}

	textLength := len(wholeText) - startsAt
	searchLength := len(startsWith)
	if searchLength > textLength {
		return false
	}

	endingLength := startsAt + searchLength
	substringFromText := wholeText[startsAt:endingLength]

	if isCaseSensitive {
		return substringFromText == startsWith
	}

	lowerCaseSubstring := strings.ToLower(substringFromText)
	lowerCaseSearchText := strings.ToLower(startsWith)

	// insensitive
	return lowerCaseSubstring == lowerCaseSearchText
}
