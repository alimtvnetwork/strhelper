package runehelper

import (
	"gitlab.com/evatix-go/strhelper/constants"
)

func IsMatch(rune1 rune, rune2 rune, isCaseSensitive bool) bool {
	if isCaseSensitive {
		return rune1 == rune2
	}

	// Insensitive case
	return ToLowerRune(rune1) == ToLowerRune(rune2)
}

func IsMatchCaseSensitive(rune1 rune, rune2 rune) bool {
	return rune1 == rune2
}

func IsMatchCaseInsensitive(rune1 rune, rune2 rune) bool {
	return ToLowerRune(rune1) == ToLowerRune(rune2)
}

func IsUpperCasePtr(r *rune) bool {
	return *r >= constants.UpperCaseA &&
		*r <= constants.UpperCaseZ
}

func IsUpperCase(r rune) bool {
	return r >= constants.UpperCaseA &&
		r <= constants.UpperCaseZ
}

func ToLowerRune(r rune) rune {
	if r >= constants.UpperCaseA &&
		r <= constants.UpperCaseZ {
		lowerCaseRune := rune(uint8(r) + constants.LowerCase)

		return lowerCaseRune
	}

	return r
}

func ToRuneArray(string *string) []rune {
	return []rune(*string)
}

func ToRuneArrayPtr(string *string) *[]rune {
	val := []rune(*string)

	return &val
}

