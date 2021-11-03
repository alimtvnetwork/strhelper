package stringsearch

import (
	"strings"
)

func IsStringsContainsSubstring(slice []string, subStringLine string) bool {
	if len(slice) == 0 {
		return false
	}

	for _, sliceItem := range slice {
		if strings.Contains(sliceItem, subStringLine) {
			return true
		}
	}

	return false
}
