package isstrs

import (
	"gitlab.com/evatix-go/strhelper/strconst"
	"gitlab.com/evatix-go/strhelper/strs/strsindex"
)

// Returns true if the findingString present in the array, if array is empty or nil then returns false.
func Exists(lines *[]string, findingString *string, isCaseSensitive bool) bool {
	return strsindex.Of(lines, findingString, 0, isCaseSensitive) > strconst.InvalidNotFoundCase
}
