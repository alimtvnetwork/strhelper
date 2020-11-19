package whitespace

import "unicode"

// Returns true for if the contents are all whitespaces
//  (including unicode whitespaces for only checking ascii use the ascii version a lot more faster)
//
// Checks from start and end if any valid char found returns immediately.
//
// Warning
//  - Panic if nil, expected to be check with nil for `runes`
//
// References:
//  - https://stackoverflow.com/a/15020162
func IsRunesWhitespaces(runes *[]rune) bool {
	// len(s) represents length in bytes so if there any unicode char it will not match with len(runes)
	length := len(*runes)
	isEven := length%2 == 0
	mid := length / 2 // 5/2 should return 2
	midLessThanOne := mid - 1
	lastIndex := length - 1
	for i := 0; i <= mid; i++ {
		rune := (*runes)[i]
		if !((rune <= maxUnit8 && asciiSpaces[rune] == 1) || (rune > maxUnit8 && unicode.IsSpace(rune))) {
			return false
		}

		if i == mid || (isEven && midLessThanOne == i) {
			// already tested above and reached the end
			break
		}

		lastIndex = lastIndex - i
		rune = (*runes)[lastIndex]

		if !((rune <= maxUnit8 && asciiSpaces[rune] == 1) || (rune > maxUnit8 && unicode.IsSpace(rune))) {
			return false
		}
	}

	return true
}
