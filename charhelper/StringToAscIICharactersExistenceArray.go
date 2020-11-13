package charhelper

import (
	"unicode"
)

// if nil or empty then return nil
// returns an 256 uint8 array for existing characters only flag to 1 which ascii characters are present in the string.
// it will reveal which chars are present
func StringToAscIICharactersExistenceArray(str *string) *[256]uint8 {
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
