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
	MaxUnit8          = 255
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
	0x85:           one, // reference : https://bit.ly/2JWdIoj
	0xA0:           one, // reference : https://bit.ly/2JWdIoj
}

// FormFeed \f is also marked as newline here.
var asciiNewLinesCharArray = [256]uint8{
	LineFeedUnix:   one,
	FormFeed:       one,
	CarriageReturn: one,
}

func GetAscIISpaceArray() [256]uint8 {
	return asciiSpace
}

// FormFeed \f is also marked as newline here.
func GetAscIINewLinesArray() [256]uint8 {
	return asciiNewLinesCharArray
}

// s == constants.EmptyString || len(s) == 0
func IsEmpty(s string) bool {
	return s == constants.EmptyString || len(s) == 0
}

func IsAscIIWhiteSpace(char uint8) bool {
	return asciiSpace[char] == 1
}

func IsUnicodeRune(char rune) bool {
	return char >= utf8.RuneSelf
}

func IsCharNotWhitespace(char uint8) bool {
	return !(asciiSpace[char] == 1 || unicode.IsSpace(rune(char)))
}

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
func IsASCIIWhitespaceOnly(s *string) bool {
	length := len(*s)
	isEven := length%2 == 0
	mid := length / 2 // 5/2 should return 2
	midLessThanOne := mid - 1
	lastIndex := length - 1
	for i := 0; i <= mid; i++ {
		char := (*s)[i]
		if !(asciiSpace[char] == 1) {
			return false
		}

		if i == mid || (isEven && midLessThanOne == i) {
			// already tested above and reached the end
			break
		}

		lastIndex = lastIndex - i
		char = (*s)[lastIndex]

		if !(asciiSpace[char] == 1) {
			return false
		}
	}

	return true
}

// Returns true for ASCII spaces and also all unicode spaces.
//
// Checks from start and end if any valid char found returns immediately.
//
//   Note: expensive operation use it wisely, needs conversion to []rune which is expensive
//
// Warning
//  - Panic if nil, expected to be check with nil for `s`
//
// References:
//  - https://stackoverflow.com/a/15020162
func IsWhitespaceOnly(s *string) bool {
	runes := []rune(*s)
	// len(s) represents length in bytes so if there any unicode char it will not match with len(runes)
	length := len(runes)
	isEven := length%2 == 0
	mid := length / 2 // 5/2 should return 2
	midLessThanOne := mid - 1
	lastIndex := length - 1
	for i := 0; i <= mid; i++ {
		rune := runes[i]
		if !((rune <= MaxUnit8 && asciiSpace[rune] == 1) || (rune > MaxUnit8 && unicode.IsSpace(rune))) {
			return false
		}

		if i == mid || (isEven && midLessThanOne == i) {
			// already tested above and reached the end
			break
		}

		lastIndex = lastIndex - i
		rune = runes[lastIndex]

		if !((rune <= MaxUnit8 && asciiSpace[rune] == 1) || (rune > MaxUnit8 && unicode.IsSpace(rune))) {
			return false
		}
	}

	return true
}

// Returns the whitespace (including unicode whitespaces) counts
//
// Note: Since unicode is in the calculation, it requires str to be in ([]rune) format which requires more memory and cost.
func AllWhitespaceCount(s *string, startsAt int) int {
	if s == nil || len(*s) == 0 {
		return 0
	}

	length := len(*s)

	if startsAt < 0 || length-1 < startsAt {
		panic("startsAt could be larger than length or less then 0.")
	}

	spacesFound := 0
	allRunes := []rune(*s)
	// needs to reset again since length may change depending on unicode chars
	// whenever rune is in the calculation it needs to be recalculated.
	length = len(allRunes)
	var rune rune
	for ; startsAt < length; startsAt++ {
		rune = allRunes[startsAt]
		if (rune <= MaxUnit8 && asciiSpace[rune] == 1) || (rune > MaxUnit8 && unicode.IsSpace(rune)) {
			spacesFound++
		}
	}

	return spacesFound
}

// Returns the whitespace (excluding unicode whitespaces) counts
func AllASCIIWhitespaceOnlyCount(s *string, startsAt int) int {
	if s == nil || len(*s) == 0 {
		return 0
	}

	length := len(*s)

	if startsAt < 0 || length-1 < startsAt {
		panic("startsAt could be larger than length or less then 0.")
	}

	spacesFound := 0
	var char uint8
	for ; startsAt < length; startsAt++ {
		char = (*s)[startsAt]
		if asciiSpace[char] == 1 {
			spacesFound++
		}
	}

	return spacesFound
}

func AllNewLinesCount(s *string, startsAt int) int {
	if s == nil || len(*s) == 0 {
		return 0
	}

	length := len(*s)

	if startsAt < 0 || length-1 < startsAt {
		panic("startsAt could be larger than length or less then 0.")
	}

	newLineFound := 0
	allRunes := []rune(*s)
	length = len(allRunes)
	var rune rune

	for ; startsAt < length; startsAt++ {
		rune = allRunes[startsAt]
		if rune <= MaxUnit8 && asciiNewLinesCharArray[rune] == 1 {
			newLineFound++
		}
	}

	return newLineFound
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

// returns s == nil || *s == constants.EmptyString || len(*s) == 0 || IsASCIIWhitespaceOnly(s)
// Checks only asc whitespaces, return false for any unicode whitespace
func IsBlankASCII(s string) bool {
	return  s == constants.EmptyString || len(s) == 0 || IsASCIIWhitespaceOnly(&s)
}

// returns s == nil || *s == constants.EmptyString || len(*s) == 0 || IsASCIIWhitespaceOnly(s)
// Checks only asc whitespaces, return false for any unicode whitespace
func IsBlankASCIIPtr(s *string) bool {
	return s == nil || *s == constants.EmptyString || len(*s) == 0 || IsASCIIWhitespaceOnly(s)
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
