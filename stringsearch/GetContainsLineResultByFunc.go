package stringsearch

import (
	"gitlab.com/auk-go/strhelper/strhelpercore"
)

func GetContainsLineResultByFunc(
	contentsLines []string,
	isLineContainsFunc IsLineContainsFunc,
) *strhelpercore.StringResult {
	if len(contentsLines) == 0 {
		return strhelpercore.InvalidStringResult()
	}

	return GetContainsLineResultByFuncPtr(
		contentsLines,
		isLineContainsFunc)
}
