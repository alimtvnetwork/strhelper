package stringsearch

import (
	"strings"

	"gitlab.com/evatix-go/core/coredata/corestr"
	"gitlab.com/evatix-go/strhelper/strhelpercore"
)

func GetContainsLineResultPtr(
	contentsLines []string,
	substrSearchLine string,
) *strhelpercore.StringResult {
	if corestr.LengthOfStrings(contentsLines) == 0 {
		return strhelpercore.InvalidStringResult()
	}

	for index, currentLine := range contentsLines {
		if strings.Contains(currentLine, substrSearchLine) {
			return &strhelpercore.StringResult{
				FoundIndex: index,
				Line:       currentLine,
				IsFound:    true,
			}
		}
	}

	return strhelpercore.InvalidStringResult()
}
