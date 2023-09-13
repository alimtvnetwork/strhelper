package hasstr

import "gitlab.com/auk-go/core/constants"

// Char
//
// Has at least one character any, returns true even if a whitespace
func Char(s string) bool {
	return s != constants.EmptyString
}
