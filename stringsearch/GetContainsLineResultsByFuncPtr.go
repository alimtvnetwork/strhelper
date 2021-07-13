package stringsearch

import (
	"gitlab.com/evatix-go/core/coredata/corestr"
	"gitlab.com/evatix-go/core/defaultcapacity"
	"gitlab.com/evatix-go/strhelper/strhelpercore"
)

func GetContainsLineResultsByFuncPtr(
	contentsLines *[]string,
	isLineContainsFunc IsLineContainsFunc,
) *strhelpercore.StringResultsMap {
	length := corestr.LengthOfStrings(contentsLines)

	if length == 0 {
		return strhelpercore.EmptyStringResultsMap()
	}

	capacity := defaultcapacity.OfSearch(length)
	currentMap := strhelpercore.NewStringResultsMap(capacity)

	for index, currentLine := range *contentsLines {
		if isLineContainsFunc(index, currentLine) {
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
