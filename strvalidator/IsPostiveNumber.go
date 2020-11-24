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

// Example : https://play.golang.org/p/NohlCrIj77r
func IsAsciiPositiveNumberPtr(str *string) bool {
	if str == nil || *str == "" || (*str)[0] == '-' {
		return false
	}

	processingPointer := str
	firstChar := (*str)[0]
	isPositiveSign := firstChar == '+'

	if isPositiveSign && len(*str) <= 1 {
		return false
	}

	if isPositiveSign {
		// Reference : https://blog.golang.org/slices-intro | https://i.imgur.com/O3Hlmac.png
		// no copy just points
		newPointers := (*str)[1:]
		processingPointer = &newPointers
	}

	isSingleDotFound := false

	for _, c := range *processingPointer {
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

// Example : https://play.golang.org/p/NohlCrIj77r
func IsUnicodePositiveNumberPtr(str *string) bool {
	if str == nil || *str == "" {
		return false
	}

	allRunes := []rune(*str)

	return IsUnicodeRunesNumberPtr(&allRunes)
}

// Example : https://play.golang.org/p/NohlCrIj77r
func IsUnicodeRunesPositiveNumberPtr(allRunes *[]rune) bool {
	if allRunes == nil || *allRunes == nil || (*allRunes)[0] == '-' {
		return false
	}

	processingPointer := allRunes
	firstChar := (*allRunes)[0]
	isPositiveSign := firstChar == '+'

	if isPositiveSign && len(*allRunes) <= 1 {
		return false
	}

	if isPositiveSign {
		// Reference : https://blog.golang.org/slices-intro | https://i.imgur.com/O3Hlmac.png
		// no copy just points
		newPointers := (*allRunes)[1:]
		processingPointer = &newPointers
	}

	isSingleDotFound := false

	for _, r := range *processingPointer {
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
