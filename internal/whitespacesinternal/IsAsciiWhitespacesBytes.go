package whitespacesinternal

import "gitlab.com/evatix-go/core/constants"

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
func IsAsciiWhitespacesBytes(bytes []byte) bool {
	length := len(bytes)
	mid := length / 2 // 5/2 should return 2
	lastIndex := length - 1
	for i := 0; i <= mid; i++ {
		char := (bytes)[i]
		if !(constants.AsciiSpace[char] == 1) {
			return false
		}

		if i == mid {
			// already tested above and reached the end
			break
		}

		lastIndex = lastIndex - i
		char = bytes[lastIndex]

		if !(constants.AsciiSpace[char] == 1) {
			return false
		}
	}

	return true
}
