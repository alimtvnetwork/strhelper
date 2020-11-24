package strvalidator

import "unicode"

func IsAsciiPositiveNumber(str string) bool {
	return IsAsciiPositiveNumberPtr(&str)
}

func IsUnicodePositiveNumber(str string) bool {
	return IsUnicodePositiveNumberPtr(&str)
}

func IsUniCodeRunesPositiveNumber(allRunes []rune) bool {
	return IsUnicodeRunesPositiveNumberPtr(&allRunes)
}

// Example : https://play.golang.org/p/VXhduA3jfP4
func IsAsciiPositiveNumberPtr(str *string) bool {
	if str == nil || *str == "" || (*str)[0] == '-' {
		return false
	}

	isSingleDotFound := false

	for _, c := range *str {
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

// Example : https://play.golang.org/p/VXhduA3jfP4
func IsUnicodePositiveNumberPtr(str *string) bool {
	if str == nil || *str == "" {
		return false
	}

	allRunes := []rune(*str)

	return IsUnicodeRunesNumberPtr(&allRunes)
}

// Example : https://play.golang.org/p/VXhduA3jfP4
func IsUnicodeRunesPositiveNumberPtr(allRunes *[]rune) bool {
	if allRunes == nil || *allRunes == nil || (*allRunes)[0] == '-' {
		return false
	}

	isSingleDotFound := false

	for _, r := range *allRunes {
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
