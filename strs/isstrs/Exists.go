package isstrs

import (
	"gitlab.com/evatix-go/core/constants"

	"gitlab.com/evatix-go/strhelper/strs/strsindex"
)

// Exists
//
// Returns true if the findingString present in the array, if array is empty or nil then returns false.
func Exists(lines []string, findingString string) bool {
	return strsindex.Of(
		lines,
		findingString,
		constants.Zero,
		true) > constants.InvalidNotFoundCase
}
