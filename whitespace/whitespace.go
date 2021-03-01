package whitespace

import (
	"gitlab.com/evatix-go/core/constants"
)

// IsEmpty(s) || IsWhitespaces(&s)
func IsNullOrWhitespace(s string) bool {
	return s == constants.EmptyString || len(s) == 0 || IsWhitespaces(&s)
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

// s == nil || *s == constants.EmptyString || len(*s) == 0 || IsWhitespaces(s)
func IsNullOrWhitespacePtr(s *string) bool {
	return s == nil || *s == constants.EmptyString || len(*s) == 0 || IsWhitespaces(s)
}

// returns true if IsNullOrWhitespace(s)
func IsBlank(s string) bool {
	return s == constants.EmptyString || len(s) == 0 || IsWhitespaces(&s)
}

// returns s == nil || *s == constants.EmptyString || len(*s) == 0 || IsWhitespaces(s)
func IsBlankPtr(s *string) bool {
	return s == nil || *s == constants.EmptyString || len(*s) == 0 || IsWhitespaces(s)
}

// returns s == nil || *s == constants.EmptyString || len(*s) == 0 || IsAsciiWhitespaces(s)
// Checks only asc whitespaces, return false for any unicode whitespace
func IsBlankAscii(s string) bool {
	return s == constants.EmptyString || len(s) == 0 || IsAsciiWhitespaces(&s)
}

// returns s == nil || *s == constants.EmptyString || len(*s) == 0 || IsAsciiWhitespaces(s)
// Checks only asc whitespaces, return false for any unicode whitespace
func IsBlankAsciiPtr(s *string) bool {
	return s == nil || *s == constants.EmptyString || len(*s) == 0 || IsAsciiWhitespaces(s)
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
