package stringsearch

import (
	"gitlab.com/auk-go/core/enums/stringcompareas"
	"gitlab.com/auk-go/strhelper/strhelpercore"
)

func GetContainsLineResultByCompareMethod(
	compareAs stringcompareas.Variant,
	contentsLines []string,
	line string,
	isCaseSensitive bool,
) *strhelpercore.StringResult {
	lineCompareFunc := compareAs.IsLineCompareFunc()

	return getContainsLineResultsUsingCompareFunc(
		lineCompareFunc,
		contentsLines,
		line,
		isCaseSensitive,
	)
}
