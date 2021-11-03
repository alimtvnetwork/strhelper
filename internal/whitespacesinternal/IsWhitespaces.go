package whitespacesinternal

import (
	"unicode"

	"gitlab.com/evatix-go/core/constants"
)

// IsWhitespaces
//
// Returns true for ASCII spaces and also all unicode spaces.
//
// Checks from start and end if any valid char found returns immediately.
//
//   Note: expensive operation use it wisely, needs conversion to []rune which is expensive
//
// Warning / Unhandled cases:
//  - Panic if nil, expected to be check with nil for `s`
//
// References:
//  - https://play.golang.org/p/78uFF8s-Dw1
func IsWhitespaces(s string) bool {
	if s == "" {
		return true
	}

	runes := []rune(s)
	// len(s) represents length in bytes so
	// if there is any unicode char
	// it will not match with len(runes)
	length := len(runes)

	if length == 0 {
		return true
	}

	mid := length / 2 // 5/2 should return 2
	lastIndex := length - 1
	var r rune
	for i := 0; i <= mid; i++ {
		r = runes[i]
		if !((r <= constants.MaxUnit8Rune && constants.AsciiSpace[r] == 1) ||
			(r > constants.MaxUnit8Rune && unicode.IsSpace(r))) {
			return false
		}

		if i == mid {
			// already tested above and reached the end
			break
		}

		lastIndex -= i
		r = runes[lastIndex]
		if !((r <= constants.MaxUnit8Rune && constants.AsciiSpace[r] == 1) ||
			(r > constants.MaxUnit8Rune && unicode.IsSpace(r))) {
			return false
		}
	}

	return true
}
