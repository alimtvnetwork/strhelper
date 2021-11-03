package whitespace

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coreindexes"
)

// IsAsciiWhitespacesBytes
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
func IsAsciiWhitespacesBytes(inputBytes []byte) bool {
	length := len(inputBytes)
	if length == 1 && (inputBytes)[coreindexes.I0] == constants.SpaceChar {
		return true
	}

	// 5/2 should return 2
	mid := length / 2
	lastIndex := length - 1
	for i := 0; i <= mid; i++ {
		char := inputBytes[i]
		if !(asciiSpaces[char] == constants.One) {
			return false
		}

		if i == mid {
			// already tested above
			// and reached the end
			break
		}

		lastIndex -= i
		char = inputBytes[lastIndex]

		if !(asciiSpaces[char] == constants.One) {
			return false
		}
	}

	return true
}
