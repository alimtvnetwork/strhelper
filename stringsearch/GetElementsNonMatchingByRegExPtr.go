package stringsearch

import (
	"regexp"

	"gitlab.com/evatix-go/core/coredata/stringslice"
)

// GetElementsNonMatchingByRegExPtr returns the lines which doesn't meet with regex requirements
func GetElementsNonMatchingByRegExPtr(
	contentsLines *[]string,
	regexp *regexp.Regexp,
) *[]string {
	if stringslice.IsEmptyPtr(contentsLines) {
		return &[]string{}
	}

	results := GetElementsNonMatchingByRegEx(
		*contentsLines, regexp)

	return &results
}
