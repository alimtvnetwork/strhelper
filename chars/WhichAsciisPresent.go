package chars

import (
	"unicode"
)

// Returns an 256 uint8 array for existing characters only flag to 1 which ascii characters are present in the string.
//  It will reveal which chars are present.
// Invalid Cases (return nil):
//  - length == 0
// *[256]uint8:
//  - represents all ASCII characters in a simple array format,
//          only existing ones which are passed will be marked with 1.
func WhichAsciisPresent(str *string) *[256]uint8 {
	length := len(*str)

	if length == 0 {
		return nil
	}

	chars := [256]uint8{}

	for _, r := range *str {
		if r <= unicode.MaxLatin1 {
			chars[r] = 1
		}
	}

	return &chars
}
