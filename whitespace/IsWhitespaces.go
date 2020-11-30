package whitespace

import "unicode"

// Returns true for if the contents are all whitespaces
//  (including unicode whitespaces for only checking ascii use the ascii version a lot more faster)
//
// Checks from start and end if any valid char found returns immediately.
//
//   Note: expensive operation use it wisely, needs conversion to []rune which is expensive
//
// Warning
//  - Panic if nil, expected to be check with nil for `s`
//
// References:
//  - https://play.golang.org/p/78uFF8s-Dw1
func IsWhitespaces(s *string) bool {
	runes := []rune(*s)
	// len(s) represents length in bytes so if there any unicode char it will not match with len(runes)
	length := len(runes)
	mid := length / 2 // 5/2 should return 2
	lastIndex := length - 1
	var r rune
	for i := 0; i <= mid; i++ {
		r = runes[i]
		if !((r <= maxUnit8 && asciiSpaces[r] == 1) || (r > maxUnit8 && unicode.IsSpace(r))) {
			return false
		}

		if i == mid {
			// already tested above and reached the end
			break
		}

		lastIndex = lastIndex - i
		r = runes[lastIndex]

		if !((r <= maxUnit8 && asciiSpaces[r] == 1) || (r > maxUnit8 && unicode.IsSpace(r))) {
			return false
		}
	}

	return true
}
