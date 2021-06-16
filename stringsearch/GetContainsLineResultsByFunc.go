package stringsearch

import (
	"gitlab.com/evatix-go/strhelper/strhelpercore"
	"gitlab.com/evatix-go/strhelper/stringcompareas"
)

func GetContainsLineResultsByFunc(
	contentsLines []string,
	isLineContainsFunc stringcompareas.IsLineContainsFunc,
) *strhelpercore.StringResultsMap {
	if len(contentsLines) == 0 {
		return strhelpercore.EmptyStringResultsMap()
	}

	return GetContainsLineResultsByFuncPtr(
		&contentsLines,
		isLineContainsFunc)
}

