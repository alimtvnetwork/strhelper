package stringsearch

import (
	"gitlab.com/evatix-go/strhelper/strhelpercore"
	"gitlab.com/evatix-go/strhelper/stringcompareas"
)

func GetContainsLineResultsMapByCompareMethod(
	compareAs stringcompareas.Variant,
	contentsLines *[]string,
	line string,
	isCaseSensitive bool,
) *strhelpercore.StringResultsMap {
	lineCompareFunc := compareAs.IsLineCompareFunc()

	return getContainsLineResultsMapUsingCompareFunc(
		lineCompareFunc,
		contentsLines,
		line,
		isCaseSensitive,
	)
}
