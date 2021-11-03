package stringsearch

import (
	"gitlab.com/evatix-go/strhelper/strhelpercore"
)

func GetEqualLineResult(
	contentsLines []string,
	line string,
) *strhelpercore.StringResult {
	if len(contentsLines) == 0 {
		return strhelpercore.InvalidStringResult()
	}

	for index, currentLine := range contentsLines {
		if currentLine == line {
			return &strhelpercore.StringResult{
				FoundIndex: index,
				Line:       currentLine,
				IsFound:    true,
			}
		}
	}

	return strhelpercore.InvalidStringResult()
}
