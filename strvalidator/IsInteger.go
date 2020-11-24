package strvalidator

import "unicode"

func IsAsciiIntegerNumber(str string) bool {
	return IsAsciiIntegerNumberPtr(&str)
}

func IsUnicodeIntegerNumber(str string) bool {
	return IsUnicodeIntegerNumberPtr(&str)
}

func IsUniCodeRunesIntegerNumber(allRunes []rune) bool {
	return IsUnicodeRunesIntegerNumberPtr(&allRunes)
}

func IsAsciiIntegerNumberPtr(str *string) bool {
	if str == nil || *str == "" {
		return false
	}
	firstChar := (*str)[0]
	isFirstCharSign := firstChar == '-' || firstChar == '+'

	if isFirstCharSign {
		if len(*str) <=1 {
			return false
		}

		for _, c := range (*str)[1:] {
			if !('0' <= c && c <= '9') {
				return false
			}
		}
	} else {
		for _, c := range *str {
			if !('0' <= c && c <= '9') {
				return false
			}
		}
	}

	return true
}

func IsUnicodeIntegerNumberPtr(str *string) bool {
	if str == nil || *str == "" {
		return false
	}

	allRunes := []rune(*str)

	return IsUnicodeRunesIntegerNumberPtr(&allRunes)
}

func IsUnicodeRunesIntegerNumberPtr(allRunes *[]rune) bool {
	if allRunes == nil || *allRunes == nil {
		return false
	}

	firstRune := (*allRunes)[0]
	isFirstCharSign := firstRune == '-' || firstRune == '+'

	if isFirstCharSign {
		if len(*allRunes) <=1 {
			return false
		}

		for _, r := range (*allRunes)[1:] {
			if !unicode.IsDigit(r) {
				return false
			}
		}
	} else {
		for _, r := range *allRunes {
			if !unicode.IsDigit(r) {
				return false
			}
		}
	}

	return true
}
