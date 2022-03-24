package stringsearch

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/converters"
)

// GetMissingElements returns elements from findingElements which doesn't contain in slice.
func GetMissingElements(slice, findElements []string) []string {
	if len(slice) == 0 || len(findElements) == 0 {
		return []string{}
	}

	missingElements := make(
		[]string,
		constants.Zero,
		len(findElements))

	hashset := converters.StringsTo.Hashset(slice)

	for _, element := range findElements {
		_, has := hashset[element]

		if !has {
			missingElements = append(
				missingElements,
				element)
		}
	}

	return missingElements
}
