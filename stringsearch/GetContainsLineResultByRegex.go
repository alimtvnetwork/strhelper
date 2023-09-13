package stringsearch

import (
	"regexp"

	"gitlab.com/auk-go/strhelper/strhelpercore"
)

func GetContainsLineResultByRegex(
	contentsLines []string,
	regexp *regexp.Regexp,
) *strhelpercore.StringResult {
	if len(contentsLines) == 0 {
		return strhelpercore.InvalidStringResult()
	}

	for index, currentLine := range contentsLines {
		if regexp.MatchString(currentLine) {
			return &strhelpercore.StringResult{
				FoundIndex: index,
				Line:       currentLine,
				IsFound:    true,
			}
		}
	}

	return strhelpercore.InvalidStringResult()
}
