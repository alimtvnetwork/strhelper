package runehelper

import (
	"gitlab.com/evatix-go/strhelper/constants"
)

func IsEitherOneNull(rune1 *rune, rune2 *rune) bool {
	return (rune1 == nil && rune2 != nil) ||
		(rune2 == nil && rune1 != nil)
}

func IsMatchPtr(rune1 *rune, rune2 *rune, isCaseSensitive bool) bool {
	isBothNull := rune1 == nil && rune2 == nil

	if isBothNull {
		return true
	}

	isEitherOneNil := IsEitherOneNull(rune1, rune2)

	if isCaseSensitive {
		return !isEitherOneNil && *rune1 == *rune2
	}

	// Insensitive case
	return !isEitherOneNil && isMatchCaseInsensitivePtr(rune1, rune2)
}

func IsMatchCaseSensitive(rune1 *rune, rune2 *rune) bool {
	isBothNull := rune1 == nil && rune2 == nil

	if isBothNull {
		return true
	}

	isEitherOneNull := IsEitherOneNull(rune1, rune2)

	return !isEitherOneNull && *rune1 == *rune2
}

func IsMatchCaseInsensitive(rune1 *rune, rune2 *rune) bool {
	isBothNull := rune1 == nil && rune2 == nil

	if isBothNull {
		return true
	}

	isEitherOneNil := IsEitherOneNull(rune1, rune2)

	return !isEitherOneNil && isMatchCaseInsensitivePtr(rune1, rune2)
}

func IsUpperCasePtr(r *rune) bool {
	return *r >= constants.UpperCaseA &&
		*r <= constants.UpperCaseZ
}

func IsUpperCase(r rune) bool {
	return r >= constants.UpperCaseA &&
		r <= constants.UpperCaseZ
}

func ToLowerRune(r *rune) *rune {
	if *r >= constants.UpperCaseA &&
		*r <= constants.UpperCaseZ {
		lowerCaseRune := rune(uint8(*r) + constants.LowerCase)

		return &lowerCaseRune
	}

	return r
}

func isMatchCaseInsensitivePtr(rune1 *rune, rune2 *rune) bool {
	return ToLowerRune(rune1) == ToLowerRune(rune2)
}

func ToRuneArray(string *string) []rune {
	return []rune(*string)
}

func ToRuneArrayPtr(string *string) *[]rune {
	val := []rune(*string)

	return &val
}

