package strs

import (
	"gitlab.com/evatix-go/strhelper/strconst"
)

// Returns true if the findingString present in the array, if array is empty or nil then returns false.
func IsContains(lines *[]string, findingString *string, isCaseSensitive bool) bool {
	return IndexOf(lines, findingString, 0, isCaseSensitive) > strconst.InvalidNotFoundCase
}
