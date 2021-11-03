package whitespace

import "gitlab.com/evatix-go/core/constants"

// IsWhitespaces Returns true for if the contents are all whitespaces
//  (including unicode whitespaces for only checking ascii use the ascii version a lot faster)
//
// Checks from start and end if any valid char found returns immediately.
//
//   Note: expensive operation use it wisely, needs conversion to []rune which is expensive
//
// Warning
//  - Panic if nil, expected to be checked with nil for `s`
//
// References:
//  - https://play.golang.org/p/78uFF8s-Dw1
func IsWhitespaces(s string) bool {
	if s == constants.EmptyString || s == doubleSpace || s == tripleSpace {
		return true
	}

	runes := []rune(s)

	return IsRunesWhitespaces(runes)
}
