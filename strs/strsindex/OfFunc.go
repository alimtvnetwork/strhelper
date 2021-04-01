package strsindex

import (
	"gitlab.com/evatix-go/core/constants"

	"gitlab.com/evatix-go/strhelper/internal/isstrsinternal"
)

// Returns the index where the string first found, rest don't care
func OfFunc(
	linesPtr *[]string,
	f func(index int, in string) (isFound, isBreak bool),
) int {
	if isstrsinternal.EmptyPtr(linesPtr) {
		return constants.InvalidNotFoundCase
	}

	for i, s := range *linesPtr {
		isFound, isBreak := f(i, s)
		if isFound {
			return i
		}

		if isBreak {
			return constants.InvalidNotFoundCase
		}
	}

	return constants.InvalidNotFoundCase
}
