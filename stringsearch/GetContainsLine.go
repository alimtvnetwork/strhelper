package stringsearch

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/corestr"
)

// GetContainsLine returns the line from the
func GetContainsLine(
	contentsLines *[]string,
	searchSubStringLine string,
) (foundLine string) {
	if corestr.LengthOfStrings(contentsLines) == 0 {
		return constants.EmptyString
	}

	for _, currentLine := range *contentsLines {
		if strings.Contains(currentLine, searchSubStringLine) {
			return currentLine
		}
	}

	return constants.EmptyString
}
