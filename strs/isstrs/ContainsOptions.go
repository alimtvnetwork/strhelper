package isstrs

import (
	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/strhelper/strs/strsindex"
)

// ContainsOptions
//
//  returns true if the findingString present in the array,
//  if array is empty or nil then returns false.
//
// One can use Exists similar to contains has fewer arguments
func ContainsOptions(
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
