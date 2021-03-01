package chars

import (
	"unicode"

	"gitlab.com/evatix-go/core/constants"
)

// *[256]uint8:
//  - represents all ASCII characters in a simple array format,
//      only existing ones which are passed will be marked with 1.
type AsciiCharacters struct {
	// asciiChars[index] == 1 meaning char exist 0 means not exist
	asciiChars *[256]uint8
}

func NewAsciiCharactersUsingString(str *string) AsciiCharacters {
	return AsciiCharacters{asciiChars: WhichAsciisPresent(str)}
}

func NewAsciiCharacters(chars *[256]uint8) AsciiCharacters {
	return AsciiCharacters{asciiChars: chars}
}

// only set if rune < 255
func (asciiCharacters *AsciiCharacters) SetRune(char rune) {
	if char <= unicode.MaxLatin1 {
		// only set if rune < 255
		(*asciiCharacters.asciiChars)[char] = 1
	}
}

func (asciiCharacters *AsciiCharacters) SetChar(char uint8) {
	(*asciiCharacters.asciiChars)[char] = 1
}

func (asciiCharacters *AsciiCharacters) UnsetChar(char uint8) {
	(*asciiCharacters.asciiChars)[char] = 0
}

func (asciiCharacters *AsciiCharacters) IsCharExists(char uint8) bool {
	return (*asciiCharacters.asciiChars)[char] == 1
}

func (asciiCharacters *AsciiCharacters) IsAnyExists(chars ...uint8) bool {
	for _, char := range chars {
		if (*asciiCharacters.asciiChars)[char] == 1 {
			return true
		}
	}

	return false
}

func (asciiCharacters *AsciiCharacters) IsAllCharsExistInStrings(strings ...*string) bool {
	for _, str := range strings {
		for _, r := range *str {
			if r > constants.MaxUnit8Rune ||
				r <= constants.MaxUnit8Rune &&
					(*asciiCharacters.asciiChars)[r] == 0 {
				return false
			}
		}
	}

	return true
}

func (asciiCharacters *AsciiCharacters) ToRunes() *[]rune {
	return AsciiArrayToRunes(asciiCharacters.asciiChars)
}

func (asciiCharacters *AsciiCharacters) String() string {
	return AsciiArrayToString(asciiCharacters.asciiChars)
}

func (asciiCharacters *AsciiCharacters) AsIs() [256]uint8 {
	return *asciiCharacters.asciiChars
}
