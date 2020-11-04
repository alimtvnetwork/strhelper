package charhelper

import (
	"gitlab.com/evatix-go/strhelper/constants"
)

func IsEitherOneNull(char1 *uint8, char2 *uint8) bool {
	return (char1 == nil && char2 != nil) ||
		(char2 == nil && char1 != nil)
}

func IsMatch(char1 uint8, char2 uint8, isCaseSensitive bool) bool {
	if isCaseSensitive {
		return char1 == char2
	}

	// Insensitive case
	return IsMatchCaseInsensitive(char1, char2)
}

func IsMatchPtr(char1 *uint8, char2 *uint8, isCaseSensitive bool) bool {
	if isCaseSensitive {
		return *char1 == *char2
	}

	// Insensitive case
	return IsMatchCaseInsensitivePtr(char1, char2)
}

func IsMatchCaseInsensitive(char1 uint8, char2 uint8) bool {
	return ToLower(char1) == ToLower(char2)
}

func IsMatchCaseInsensitivePtr(char1 *uint8, char2 *uint8) bool {
	return ToLowerPtr(char1) == ToLowerPtr(char2)
}

func IsUpperCase(c uint8) bool {
	return c >= constants.UpperCaseA &&
		c <= constants.UpperCaseZ
}

func ToLower(c uint8) uint8 {
	if c >= constants.UpperCaseA &&
		c <= constants.UpperCaseZ {
		return c + constants.LowerCase
	}

	return c
}

func IsUpperCasePtr(c *uint8) bool {
	return *c >= constants.UpperCaseA &&
		*c <= constants.UpperCaseZ
}

func ToLowerPtr(c *uint8) uint8 {
	if *c >= constants.UpperCaseA &&
		*c <= constants.UpperCaseZ {
		return *c + constants.LowerCase
	}

	return *c
}
