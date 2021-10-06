package stringsearch

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/corestr"
)

// GetContainsLinePtr returns the line from the search
//
// Returns the first item that contains the substring.
func GetContainsLinePtr(
	contentsLines []string,
	searchSubStringLine string,
) (foundLine string) {
	if corestr.LengthOfStrings(contentsLines) == 0 {
		return constants.EmptyString
	}

	for _, currentLine := range contentsLines {
		if strings.Contains(currentLine, searchSubStringLine) {
			return currentLine
		}
	}

	return constants.EmptyString
}
