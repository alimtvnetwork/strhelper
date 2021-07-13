package stringsearch

import (
	"gitlab.com/evatix-go/core/coredata/corestr"
	"gitlab.com/evatix-go/strhelper/strhelpercore"
)

func GetContainsLineResultByFuncPtr(
	contentsLines *[]string,
	isLineContainsFunc IsLineContainsFunc,
) *strhelpercore.StringResult {
	if corestr.LengthOfStrings(contentsLines) == 0 {
		return strhelpercore.InvalidStringResult()
	}

	for index, currentLine := range *contentsLines {
		if isLineContainsFunc(index, currentLine) {
			return &strhelpercore.StringResult{
				FoundIndex: index,
				Line:       currentLine,
				IsFound:    true,
			}
		}
	}

	return strhelpercore.InvalidStringResult()
}
