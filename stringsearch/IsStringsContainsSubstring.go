package stringsearch

import (
	"strings"

	"gitlab.com/evatix-go/core/coredata/corestr"
)

func IsStringsContainsSubstring(slice []string, subStringLine string) bool {
	if corestr.LengthOfStrings(slice) == 0 {
		return false
	}

	for _, sliceItem := range slice {
		if strings.Contains(sliceItem, subStringLine) {
			return true
		}
	}

	return false
}
