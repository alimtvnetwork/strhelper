package strvalidator

import "unicode"

func IsAsciiNegativeNumber(str string) bool {
	return IsAsciiNegativeNumberPtr(&str)
}

func IsUnicodeNegativeNumber(str string) bool {
	return IsUnicodeNegativeNumberPtr(&str)
}

func IsUniCodeRunesNegativeNumber(allRunes []rune) bool {
	return IsUnicodeRunesNegativeNumberPtr(&allRunes)
}

// Example : https://play.golang.org/p/UtTytkdk3KP
func IsAsciiNegativeNumberPtr(str *string) bool {
	if str == nil || *str == "" || (*str)[0] != '-' {
		return false
	}

	isSingleDotFound := false

	for _, c := range (*str)[1:] {
		if !isSingleDotFound && c == '.' {
			isSingleDotFound = true
			continue
		}

		if !(('0' <= c && c <= '9') || (!isSingleDotFound && c == '.')) {
			return false
		}
	}

	return true
}

// Example : https://play.golang.org/p/UtTytkdk3KP
func IsUnicodeNegativeNumberPtr(str *string) bool {
	if str == nil || *str == "" {
		return false
	}

	allRunes := []rune(*str)

	return IsUnicodeRunesNegativeNumberPtr(&allRunes)
}

// Example : https://play.golang.org/p/UtTytkdk3KP
func IsUnicodeRunesNegativeNumberPtr(allRunes *[]rune) bool {
	if allRunes == nil || *allRunes == nil || (*allRunes)[0] != '-' {
		return false
	}

	isSingleDotFound := false

	for _, r := range (*allRunes)[1:] {
		if !isSingleDotFound && r == '.' {
			isSingleDotFound = true
			continue
		}

		if !(unicode.IsDigit(r) || (!isSingleDotFound && r == '.')) {
			return false
		}
	}

	return true
}
