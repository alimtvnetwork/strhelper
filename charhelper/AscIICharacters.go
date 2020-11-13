package charhelper

import "unicode"

type AscIICharacters struct {
	// asciiChars[index] == 1 meaning char exist 0 means not exist
	asciiChars *[256]uint8
}

func NewAscIICharactersUsingString(str *string) AscIICharacters {
	return AscIICharacters{asciiChars: StringToAscIICharactersExistenceArray(str)}
}

func NewAscIICharacters(chars *[256]uint8) AscIICharacters {
	return AscIICharacters{asciiChars: chars}
}

// only set if rune < 255
func (ascIICharacters *AscIICharacters) SetRune(char rune) {
	if char <= unicode.MaxLatin1 {
		// only set if rune < 255
		(*(*ascIICharacters).asciiChars)[char] = 1
	}
}

func (ascIICharacters *AscIICharacters) SetChar(char uint8) {
	(*(*ascIICharacters).asciiChars)[char] = 1
}

func (ascIICharacters *AscIICharacters) UnsetChar(char uint8) {
	(*(*ascIICharacters).asciiChars)[char] = 0
}

func (ascIICharacters *AscIICharacters) IsCharExists(char uint8) bool {
	return (*(*ascIICharacters).asciiChars)[char] == 1
}

func (ascIICharacters *AscIICharacters) IsAnyExists(chars ...uint8) bool {
	for _, char := range chars {
		if (*(*ascIICharacters).asciiChars)[char] == 1 {
			return true
		}
	}

	return false
}

func (ascIICharacters *AscIICharacters) IsAllCharsExistInStrings(strings ...*string) bool {
	for _, str := range strings {
		for i := range *str {
			char := (*str)[i]

			if (*ascIICharacters.asciiChars)[char] == 0 {
				return false
			}
		}
	}

	return true
}

func (ascIICharacters *AscIICharacters) ToRunes() *[]rune {
	return AscIIArrayToRunes(ascIICharacters.asciiChars)
}

func (ascIICharacters *AscIICharacters) String() string {
	return AscIIArrayToString(ascIICharacters.asciiChars)
}

func (ascIICharacters *AscIICharacters) AsIs() [256]uint8 {
	return *ascIICharacters.asciiChars
}
