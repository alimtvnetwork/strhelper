package whitespace

import (
	"unicode"
	"unicode/utf8"

	"gitlab.com/evatix-go/strhelper/constants"
)

// Reference :
// - https://en.wikipedia.org/wiki/Newline,
// - https://en.wikipedia.org/wiki/Whitespace_character
// - https://en.wikipedia.org/wiki/Regular_expression#Character_classes

const (
	Space             = ' '
	Tab               = '\t'
	LineFeedUnix      = '\n'
	CarriageReturn    = '\r'
	FormFeed          = '\f'
	TabV              = '\v'
	LineFeedStr       = constants.NewLine
	LineFeedUnixStr   = string(LineFeedUnix)
	SpaceStr          = string(Space)
	TabStr            = string(Tab)
	CarriageReturnStr = string(CarriageReturn)
	FormFeedStr       = string(FormFeed)
	TabVStr           = string(TabV)
	one               = 1
)

// Copied from golang strings
var asciiSpace = [256]uint8{
	Tab:            one,
	LineFeedUnix:   one,
	TabV:           one,
	FormFeed:       one,
	CarriageReturn: one,
	Space:          one,
}

// s == constants.EmptyString || len(s) == 0
func IsEmpty(s string) bool {
	return s == constants.EmptyString || len(s) == 0
}

func IsNullOrAscIIWhiteSpace(char *uint8) bool {
	return char == nil || asciiSpace[*char] == 1
}

func IsUnicodeCharPtr(char *uint8) bool {
	return char != nil && rune(*char) >= utf8.RuneSelf
}

func IsUnicodeChar(char uint8) bool {
	return rune(char) >= utf8.RuneSelf
}

func IsUnicodeRune(char rune) bool {
	return char >= utf8.RuneSelf
}

func IsCharNotWhitespace(char *uint8) bool {
	return !IsNullOrAscIIWhiteSpace(char) && !unicode.IsSpace(rune(*char))
}

// Warning : Panic if nil, expected to be check with nil for `s`
func IsWhitespaceOnly(s *string) bool {
	length := len(*s)

	mid := length / 2 // 5/2 should return 2
	lastIndex := length - 1
	for i := 0; i <= mid; i++ {
		char := (*s)[i]
		if !IsNullOrAscIIWhiteSpace(&char) && !unicode.IsSpace(rune(char)) {
			return false
		}

		if i == mid {
			// already tested above and reached the end
			break
		}

		lastIndex = lastIndex - i
		char = (*s)[lastIndex]

		if !IsNullOrAscIIWhiteSpace(&char) && !unicode.IsSpace(rune(char)) {
			return false
		}
	}

	return true
}

// IsEmpty(s) || IsWhitespaceOnly(&s)
func IsNullOrWhitespace(s string) bool {
	return s == constants.EmptyString || len(s) == 0 || IsWhitespaceOnly(&s)
}

// Has at least one character any, returns true even if a whitespace
func HasCharacter(s string) bool {
	return !(s == constants.EmptyString || len(s) == 0)
}

// Has at least one character any, returns true even if a whitespace
func HasCharacterPtr(s *string) bool {
	return !(s == nil || *s == constants.EmptyString || len(*s) == 0)
}

// Has at least one character other than space or whitespace
func HasCharacterWithoutWhitespaces(s string) bool {
	return !(s == constants.EmptyString || len(s) == 0 || IsNullOrWhitespacePtr(&s))
}

// Has at least one character other than space or whitespace
func IsDefinedWithCharsWithoutWhitespaces(s string) bool {
	return !(s == constants.EmptyString || len(s) == 0 || IsNullOrWhitespacePtr(&s))
}

// Has at least one character other than space or whitespace
func HasCharacterWithoutWhitespacesPtr(s *string) bool {
	return !(s == nil || *s == constants.EmptyString || len(*s) == 0 || IsNullOrWhitespacePtr(s))
}

// Has at least one character other than space or whitespace
func IsDefinedWithCharsWithoutWhitespacesPtr(s *string) bool {
	return !(s == nil || *s == constants.EmptyString || len(*s) == 0 || IsNullOrWhitespacePtr(s))
}

// s == nil || *s == constants.EmptyString || len(*s) == 0 || IsWhitespaceOnly(s)
func IsNullOrWhitespacePtr(s *string) bool {
	return s == nil || *s == constants.EmptyString || len(*s) == 0 || IsWhitespaceOnly(s)
}

// returns true if IsNullOrWhitespace(s)
func IsBlank(s string) bool {
	return s == constants.EmptyString || len(s) == 0 || IsWhitespaceOnly(&s)
}

// returns s == nil || *s == constants.EmptyString || len(*s) == 0 || IsWhitespaceOnly(s)
func IsBlankPtr(s *string) bool {
	return s == nil || *s == constants.EmptyString || len(*s) == 0 || IsWhitespaceOnly(s)
}

// returns true if Any of the strings is blanks thus empty or whitespace or nil
func HasAnyBlank(strings ...*string) bool {
	for _, str := range strings {
		if IsBlankPtr(str) {
			return true
		}
	}

	return false
}

// returns true if All of the strings are blank thus empty or whitespaces or null/nil.
func HasAllBlanks(strings ...*string) bool {
	for _, str := range strings {
		if IsDefinedWithCharsWithoutWhitespacesPtr(str) {
			return false
		}
	}

	return true
}

// returns true if Any strings are defined thus has character other than whitespace/empty/nil
func HasAnyDefined(strings ...*string) bool {
	for _, str := range strings {
		if IsDefinedWithCharsWithoutWhitespacesPtr(str) {
			return true
		}
	}

	return false
}

// returns true if All strings are defined thus has character other than whitespace/empty/nil
func HasAllDefined(strings ...*string) bool {
	for _, str := range strings {
		if IsBlankPtr(str) {
			return false
		}
	}

	return true
}

// returns true if any of the strings is blank thus empty or whitespaces or null/nil.
// returns true for nil or empty array
func HasAnyBlankArray(strings *[]*string) bool {
	if isEmptyStringPtrArray(strings) {
		return true
	}

	for _, str := range *strings {
		if IsBlankPtr(str) {
			return true
		}
	}

	return false
}

// returns true if All of strings are blanks thus empty or whitespaces or null/nil.
// returns true for nil or empty array
func HasAllBlanksArray(strings *[]*string) bool {
	if isEmptyStringPtrArray(strings) {
		return true
	}

	for _, str := range *strings {
		if IsDefinedWithCharsWithoutWhitespacesPtr(str) {
			return false
		}
	}

	return true
}

// returns true if any of it is defined thus has character other than whitespace
// returns false for nil or empty array
func HasAnyDefinedArray(strings *[]*string) bool {
	if isEmptyStringPtrArray(strings) {
		return false
	}

	for _, str := range *strings {
		if IsDefinedWithCharsWithoutWhitespacesPtr(str) {
			return true
		}
	}

	return false
}

// returns true if All of it is defined thus has character other than whitespace
// returns false for nil or empty array
func HasAllDefinedArray(strings *[]*string) bool {
	if isEmptyStringPtrArray(strings) {
		return false
	}

	for _, str := range *strings {
		if IsBlankPtr(str) {
			return false
		}
	}

	return true
}
