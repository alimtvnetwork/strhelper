package splits

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/corestr"
)

// EachLineSplit
// Issue : https://gitlab.com/evatix-go/strhelper/-/issues/118
// Split whole text by new line first and then each line split by the given eachLineSplitBy
// Left is the key and right is the value
func EachLineSplit(wholeText, eachLineSplitBy string) []corestr.KeyValuePair {
	lines := strings.Split(wholeText, eachLineSplitBy)

	slice := make([]corestr.KeyValuePair, constants.Zero, len(lines))

	for _, line := range lines {
		key, value := IntoTwo(line, eachLineSplitBy)
		keyVal := corestr.KeyValuePair{
			Key:   key,
			Value: value,
		}

		slice = append(slice, keyVal)
	}

	return slice
}
