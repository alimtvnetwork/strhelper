package stringsearch

import (
	"gitlab.com/evatix-go/strhelper/strhelpercore"
)

func GetContainsLineResultsByFunc(
	contentsLines []string,
	isLineContainsFunc IsLineContainsFunc,
) *strhelpercore.StringResultsMap {
	if len(contentsLines) == 0 {
		return strhelpercore.EmptyStringResultsMap()
	}

	return GetContainsLineResultsByFuncPtr(
		&contentsLines,
		isLineContainsFunc)
}
