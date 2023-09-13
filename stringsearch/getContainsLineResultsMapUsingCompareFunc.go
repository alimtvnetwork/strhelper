package stringsearch

import (
	"gitlab.com/auk-go/core/defaultcapacity"
	"gitlab.com/auk-go/core/enums/stringcompareas"
	"gitlab.com/auk-go/strhelper/strhelpercore"
)

func getContainsLineResultsMapUsingCompareFunc(
	isLineCompareFunc stringcompareas.IsLineCompareFunc,
	contentsLines []string,
	line string,
	isCaseSensitive bool,
) *strhelpercore.StringResultsMap {
	length := len(contentsLines)
	if len(contentsLines) == 0 {
		return strhelpercore.EmptyStringResultsMap()
	}

	capacity := defaultcapacity.OfSearch(length)
	currentMap := strhelpercore.NewStringResultsMap(capacity)

	for index, currentLine := range contentsLines {
		if isLineCompareFunc(currentLine, line, isCaseSensitive) {
			result := &strhelpercore.StringResult{
				FoundIndex: index,
				Line:       currentLine,
				IsFound:    true,
			}

			currentMap.AddFoundOnly(result)
		}
	}

	return currentMap
}
