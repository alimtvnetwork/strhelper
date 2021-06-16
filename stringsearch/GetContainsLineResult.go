package stringsearch

import (
	"gitlab.com/evatix-go/strhelper/strhelpercore"
)

func GetContainsLineResult(
	contentsLines []string,
	substrSearchLine string,
) *strhelpercore.StringResult {
	if len(contentsLines) == 0 {
		return strhelpercore.InvalidStringResult()
	}

	return GetContainsLineResultPtr(
		&contentsLines,
		substrSearchLine)
}
