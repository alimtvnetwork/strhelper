package stringsearch

import (
	"gitlab.com/evatix-go/core/coredata/corestr"
)

func IsStringsContainsEqualLine(slice *[]string, line string) bool {
	if corestr.LengthOfStrings(slice) == 0 {
		return false
	}

	for _, sliceItem := range *slice {
		if line == sliceItem {
			return true
		}
	}

	return false
}
