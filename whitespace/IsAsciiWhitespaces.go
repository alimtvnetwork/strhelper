package whitespace

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
//  - https://stackoverflow.com/a/15020162
func IsAsciiWhitespaces(s *string) bool {
	length := len(*s)
	isEven := length%2 == 0
	mid := length / 2 // 5/2 should return 2
	midLessThanOne := mid - 1
	lastIndex := length - 1
	for i := 0; i <= mid; i++ {
		char := (*s)[i]
		if !(asciiSpaces[char] == 1) {
			return false
		}

		if i == mid || (isEven && midLessThanOne == i) {
			// already tested above and reached the end
			break
		}

		lastIndex = lastIndex - i
		char = (*s)[lastIndex]

		if !(asciiSpaces[char] == 1) {
			return false
		}
	}

	return true
}
