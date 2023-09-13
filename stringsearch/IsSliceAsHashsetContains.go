package stringsearch

import (
	"gitlab.com/auk-go/core/converters"
)

func IsSliceAsHashsetContains(
	sliceAsHashset []string,
	containsLine string,
) bool {
	if len(sliceAsHashset) == 0 {
		return false
	}

	hashset := converters.StringsTo.Hashset(sliceAsHashset)
	_, has := hashset[containsLine]

	return has
}
