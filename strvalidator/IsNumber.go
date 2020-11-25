package strvalidator

import "unicode"

func IsAsciiNumber(str string) bool {
	return IsAsciiNumberPtr(&str)
}

func IsUnicodeNumber(str string) bool {
	return IsUnicodeNumberPtr(&str)
}

func IsUniCodeRunesNumber(allRunes []rune) bool {
	return IsUnicodeRunesNumberPtr(&allRunes)
}

// Example : https://play.golang.org/p/_314bvo6TDk
func IsAsciiNumberPtr(str *string) bool {
	if str == nil || *str == "" {
		return false
	}

	firstChar := (*str)[0]
	isFirstCharSign := firstChar == '-' || firstChar == '+'
	isSingleDotFound := false

	if isFirstCharSign {
		if len(*str) <= 1 {
			return false
		}

		for _, c := range (*str)[1:] {
			if !isSingleDotFound && c == '.' {
				isSingleDotFound = true
				continue
			}

			if !(('0' <= c && c <= '9') || (!isSingleDotFound && c == '.')) {
				return false
			}
		}
	} else {
		for _, c := range *str {
			if !isSingleDotFound && c == '.' {
				isSingleDotFound = true
				continue
			}

			if !(('0' <= c && c <= '9') || (!isSingleDotFound && c == '.')) {
				return false
			}
		}
	}

	return true
}

// Example : https://play.golang.org/p/_314bvo6TDk
func IsUnicodeNumberPtr(str *string) bool {
	if str == nil || *str == "" {
		return false
	}

	allRunes := []rune(*str)

	return IsUnicodeRunesNumberPtr(&allRunes)
}

// Example : https://play.golang.org/p/_314bvo6TDk
func IsUnicodeRunesNumberPtr(allRunes *[]rune) bool {
	if allRunes == nil || *allRunes == nil {
		return false
	}

	firstRune := (*allRunes)[0]
	isFirstCharSign := firstRune == '-' || firstRune == '+'

	isSingleDotFound := false
	if isFirstCharSign {
		if len(*allRunes) <= 1 {
			return false
		}

		for _, r := range (*allRunes)[1:] {
			if !isSingleDotFound && r == '.' {
				isSingleDotFound = true
				continue
			}

			if !(unicode.IsDigit(r) || (!isSingleDotFound && r == '.')) {
				return false
			}
		}
	} else {
		for _, r := range *allRunes {
			if !isSingleDotFound && r == '.' {
				isSingleDotFound = true
				continue
			}

			if !(unicode.IsDigit(r) || (!isSingleDotFound && r == '.')) {
				return false
			}
		}
	}

	return true
}
