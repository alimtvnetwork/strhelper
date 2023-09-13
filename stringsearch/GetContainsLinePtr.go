package stringsearch

import (
	"strings"

	"gitlab.com/auk-go/core/constants"
)

// GetContainsLinePtr returns the line from the search
//
// Returns the first item that contains the substring.
func GetContainsLinePtr(
	contentsLines []string,
	searchSubStringLine string,
) (foundLine string) {
	if len(contentsLines) == 0 {
		return constants.EmptyString
	}

	for _, currentLine := range contentsLines {
		if strings.Contains(currentLine, searchSubStringLine) {
			return currentLine
		}
	}

	return constants.EmptyString
}
