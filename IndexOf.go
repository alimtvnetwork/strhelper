package strhelper

import (
	"strings"

	"gitlab.com/evatix-go/strhelper/constants"
)

// TODO Fix
func IndexOf(s, findingString string, startsAt int, isCaseSensitive bool) int {
	if isCaseSensitive && startsAt == 0 {
		return strings.Index(s, findingString)
	}

	// rest of the cases are not implemented properly.
	panic(constants.NotImplemented)

	return IndexOfLongestCommonSuffix(s, findingString, startsAt, isCaseSensitive)
}
