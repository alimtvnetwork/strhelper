package stringsearch

import (
	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/converters"
	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/strhelper/internal/messages"
)

// GetFirstMissingElements returns the first element from
// findingElements which doesn't contain in slice.
func GetFirstMissingElements(
	slice,
	findElements []string,
) *corestr.ValueStatus {
	if len(slice) == 0 || len(findElements) == 0 {
		return corestr.InvalidValueStatus(messages.SliceOrFindingElementsAreNilOrEmpty)
	}

	hashset := converters.StringsTo.Hashset(slice)

	for i, element := range findElements {
		_, has := hashset[element]

		if !has {
			return &corestr.ValueStatus{
				ValueValid: &corestr.ValidValue{
					Value:   element,
					IsValid: true,
					Message: constants.EmptyString,
				},
				Index: i,
			}
		}
	}

	return corestr.InvalidValueStatus(messages.SliceOrFindingElementsAreNilOrEmpty)
}
