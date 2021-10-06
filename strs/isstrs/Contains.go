package isstrs

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/strhelper/strs/strsindex"
)

// Contains Returns true if the findingString present in the array,
// if array is empty or nil then returns false.
//
// One can use Exists similar to contains has less arguments
func Contains(
	lines []string,
	findingString string,
	startsAt int,
	isCaseSensitive bool,
) bool {
	return strsindex.Of(
		lines,
		findingString,
		startsAt,
		isCaseSensitive) > constants.InvalidNotFoundCase
}
