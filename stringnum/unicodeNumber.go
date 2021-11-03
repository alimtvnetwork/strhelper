package stringnum

import "unicode"

// unicodeNumber
//
// represents verification of unicode to float numbers
type unicodeNumber struct{}

func (it *unicodeNumber) Positive(input string) bool {
	return it.Unsigned(input)
}

func (it *unicodeNumber) Negative(input string) bool {
	if input == "" || input[0] != '-' {
		return false
	}

	return it.Signed(input)
}

// Signed
//
// Example : https://play.golang.org/p/_314bvo6TDk
func (it *unicodeNumber) Signed(input string) bool {
	if input == "" || input == "+" || input == "-" {
		return false
	}

	firstRune := input[0]
	isFirstCharSign := firstRune == '-' || firstRune == '+'
	processingPointer := input
	if isFirstCharSign {
		processingPointer = input[1:]
	}

	isSingleDotFound := false
	for _, r := range processingPointer {
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

// Unsigned
//
// Example : https://play.golang.org/p/NohlCrIj77r
func (it *unicodeNumber) Unsigned(input string) bool {
	if input == "" || input == "+" || input == "-" || input[0] == '-' {
		return false
	}

	firstChar := input[0]
	isPositiveSign := firstChar == '+'
	processingPointer := input
	if isPositiveSign {
		processingPointer = input[1:]
	}

	isSingleDotFound := false
	for _, r := range processingPointer {
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

// UnsignedWithoutPlus
//
// Returns true for positive number without `+` symbol as well.
func (it *unicodeNumber) UnsignedWithoutPlus(input string) bool {
	if input == "" || input == "+" || input == "-" || input[0] == '-' || input[0] == '+' {
		return false
	}

	isSingleDotFound := false
	for _, r := range input {
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
