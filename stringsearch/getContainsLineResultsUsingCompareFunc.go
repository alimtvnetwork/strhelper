package stringsearch

import (
	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/enums/stringcompareas"
	"gitlab.com/auk-go/strhelper/strhelpercore"
)

func getContainsLineResultsUsingCompareFunc(
	isLineCompareFunc stringcompareas.IsLineCompareFunc,
	contentsLines []string,
	line string,
	isCaseSensitive bool,
) *strhelpercore.StringResult {
	if len(contentsLines) == 0 {
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
