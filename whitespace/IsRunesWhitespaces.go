package whitespace

import "unicode"

// IsRunesWhitespaces
//
// Returns true for if the contents are all whitespaces
//  (including unicode whitespaces for only checking ascii use the ascii version a lot more faster)
//
// Checks from start and end if any valid char found returns immediately.
//
// Warning
//  - Panic if nil, expected to be check with nil for `runes`
//
// References:
//  - https://play.golang.org/p/78uFF8s-Dw1
func IsRunesWhitespaces(runes []rune) bool {
	// len(s) represents length in bytes so
	// if there is any unicode char it will not match with len(runes)
	length := len(runes)
	// 5/2 should return 2
	mid := length / 2
	lastIndex := length - 1
	var r rune
	for i := 0; i <= mid; i++ {
		r = (runes)[i]
		if !((r <= maxUnit8 && asciiSpaces[r] == 1) || (r > maxUnit8 && unicode.IsSpace(r))) {
			return false
		}

		if i == mid {
			// already tested above and reached the end
			break
		}

		lastIndex = lastIndex - i
		r = (runes)[lastIndex]

		if !((r <= maxUnit8 && asciiSpaces[r] == 1) || (r > maxUnit8 && unicode.IsSpace(r))) {
			return false
		}
	}

	return true
}
