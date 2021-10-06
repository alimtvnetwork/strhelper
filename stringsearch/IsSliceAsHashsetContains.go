package stringsearch

import (
	"gitlab.com/evatix-go/core/converters"
	"gitlab.com/evatix-go/core/coredata/corestr"
)

func IsSliceAsHashsetContains(
	sliceAsHashset []string,
	containsLine string,
) bool {
	if corestr.LengthOfStrings(sliceAsHashset) == 0 {
		return false
	}

	hashset := *converters.StringsToMap(&sliceAsHashset)
	_, has := hashset[containsLine]

	return has
}
