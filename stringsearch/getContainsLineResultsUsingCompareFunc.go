package stringsearch

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/corestr"
	"gitlab.com/evatix-go/core/enums/stringcompareas"
	"gitlab.com/evatix-go/strhelper/strhelpercore"
)

func getContainsLineResultsUsingCompareFunc(
	isLineCompareFunc stringcompareas.IsLineCompareFunc,
	contentsLines []string,
	line string,
	isCaseSensitive bool,
) *strhelpercore.StringResult {
	if corestr.LengthOfStrings(contentsLines) == 0 {
		return &strhelpercore.StringResult{
			FoundIndex: constants.InvalidNotFoundCase,
			Line:       constants.EmptyString,
			IsFound:    false,
		}
	}

	for index, currentLine := range contentsLines {
		if isLineCompareFunc(currentLine, line, isCaseSensitive) {
			return &strhelpercore.StringResult{
				FoundIndex: index,
				Line:       currentLine,
				IsFound:    true,
			}
		}
	}

	return &strhelpercore.StringResult{
		FoundIndex: constants.InvalidNotFoundCase,
		Line:       constants.EmptyString,
		IsFound:    false,
	}
}
