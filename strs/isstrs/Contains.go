package isstrs

import (
	"gitlab.com/evatix-go/strhelper/strconst"
	"gitlab.com/evatix-go/strhelper/strs/strsindex"
)

// Returns true if the findingString present in the array, if array is empty or nil then returns false.
func Contains(lines *[]string, findingString *string, startsAt int, isCaseSensitive bool) bool {
	return strsindex.Of(lines, findingString, startsAt, isCaseSensitive) > strconst.InvalidNotFoundCase
}
