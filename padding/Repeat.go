package padding

import (
	"unsafe"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coreindexes"
)

func Repeat(s string, width int) string {
	if s == "" {
		return constants.EmptyString
	}

	length := len(s)

	if length == 0 {
		return constants.EmptyString
	}

	if width == 1 {
		return s
	}

	if width == 2 {
		return s + s
	}

	if width == 3 {
		return s + s + s
	}

	if width == 4 {
		return s + s + s + s
	}

	if width == 5 {
		return s + s + s + s + s
	}

	// return strings.Repeat(*s, width)

	wholeLength := length * width
	allChars := make([]byte, wholeLength)

	if length == constants.One {
		for i := 0; i < wholeLength; i++ {
			allChars[i] = s[coreindexes.I0]
		}

		return *(*string)(unsafe.Pointer(&allChars))
	}

	charIndex := 0

	for i := 0; i < wholeLength; i++ {
		allChars[i] = s[charIndex]

		charIndex++
		if charIndex == length {
			charIndex = 0
		}
	}

	return *(*string)(unsafe.Pointer(&allChars))
}
