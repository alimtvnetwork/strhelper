package whitespace

import "gitlab.com/auk-go/core/constants"

// IsAsciiWhitespaces
//
// Returns true for ASCII spaces only. Returns false for unicode whitespaces.
//
// If there is any unicode space it will count as character and return false.
//
// Checks from start and end if any valid char found returns immediately.
//
// Warning
//  - Panic if nil, expected to be check with nil for `s`
//
// References:
//  - https://play.golang.org/p/78uFF8s-Dw1
func IsAsciiWhitespaces(s string) bool {
	if s == constants.EmptyString || s == doubleSpace || s == tripleSpace {
		return true
	}

	length := len(s)
	// 5/2 should return 2
	mid := length / 2
	lastIndex := length - 1

	for i := 0; i <= mid; i++ {
		char := s[i]
		if !(asciiSpaces[char] == constants.One) {
			return false
		}

		if i == mid {
			// already tested above and reached the end
			break
		}

		lastIndex -= i
		char = s[lastIndex]

		if !(asciiSpaces[char] == constants.One) {
			return false
		}
	}

	return true
}
