package stringsearch

import (
	"gitlab.com/evatix-go/strhelper/strhelpercore"
	"gitlab.com/evatix-go/strhelper/stringcompareas"
)

func GetContainsLineResultByFunc(
	contentsLines []string,
	isLineContainsFunc stringcompareas.IsLineContainsFunc,
) *strhelpercore.StringResult {
	if len(contentsLines) == 0 {
		return strhelpercore.InvalidStringResult()
	}

	return GetContainsLineResultByFuncPtr(
		&contentsLines,
		isLineContainsFunc)
}
