package stringsearch

import (
	"gitlab.com/evatix-go/core/coredata/corestr"
	"gitlab.com/evatix-go/core/defaultcapacity"
	"gitlab.com/evatix-go/strhelper/strhelpercore"
	"gitlab.com/evatix-go/strhelper/stringcompareas"
)

func getContainsLineResultsMapUsingCompareFunc(
	isLineCompareFunc stringcompareas.IsLineCompareFunc,
	contentsLines *[]string,
	line string,
	isCaseSensitive bool,
) *strhelpercore.StringResultsMap {
	length := corestr.LengthOfStrings(contentsLines)
	if corestr.LengthOfStrings(contentsLines) == 0 {
		return strhelpercore.EmptyStringResultsMap()
	}

	capacity := defaultcapacity.OfSearch(length)
	currentMap := strhelpercore.NewStringResultsMap(capacity)

	for index, currentLine := range *contentsLines {
		if isLineCompareFunc(currentLine, line, isCaseSensitive) {
			result := &strhelpercore.StringResult{
				FoundIndex: index,
				Line:       currentLine,
				IsFound:    true,
			}

			currentMap.Add(result)
		}
	}

	return currentMap
}
