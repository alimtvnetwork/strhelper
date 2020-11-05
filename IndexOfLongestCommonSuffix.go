package strhelper

import (
	"strings"

	"gitlab.com/evatix-go/strhelper/constants"
	"gitlab.com/evatix-go/strhelper/whitespace"
)

// Code Copied from Reference: https://bit.ly/35ZGJHc
// Has Longest common suffix, returns -1 if doesn't found.
// If any is nil or has constants.EmptyString empty string then it return -1
// TODO : https://gitlab.com/evatix-go/strhelper/-/issues/2
func IndexOfLongestCommonSuffix(
	a, b string,
	startsAt int,
	isCaseSensitive bool,
) int {
	panic(constants.NotImplemented + " https://gitlab.com/evatix-go/strhelper/-/issues/2")
	if whitespace.HasAnyBlank(&a, &b) {
		return constants.InvalidNotFoundCase
	}

	lenA := len(a)
	lenB := len(b)

	if isCaseSensitive {
		// both needs to be in same case
		a = strings.ToLower(a)
		b = strings.ToLower(b)
	}

	i := startsAt
	for ; i < lenA && i < lenB; i++ {
		if a[lenA-1-i] != b[lenB-1-i] {
			return i
		}
	}

	return i
}
