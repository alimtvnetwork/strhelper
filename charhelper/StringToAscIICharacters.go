package charhelper

import (
	"unicode"
)

// if nil or empty then return nil
// returns an 256 uint8 array for existing characters only.
// it will reveal which chars are present
func StringToAscIICharacters(str *string) *[256]uint8 {
	length := len(*str)

	if length == 0 {
		return nil
	}

	chars := [256]uint8{}

	for _, char := range *str {
		if char <= unicode.MaxLatin1 {
			chars[char] = 1
		}
	}

	return &chars
}
