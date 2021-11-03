package stringnum

import "unicode"

// unicodeInteger
//
// represents verification of unicode to integer numbers
type unicodeInteger struct{}

func (it *unicodeInteger) Positive(input string) bool {
	return it.Unsigned(input)
}

func (it *unicodeInteger) Negative(input string) bool {
	if input == "" || input[0] != '-' {
		return false
	}

	return it.Signed(input)
}

// Signed
//
// Example : https://play.golang.org/p/_314bvo6TDk
func (it *unicodeInteger) Signed(input string) bool {
	if input == "" || input == "+" || input == "-" {
		return false
	}

	firstRune := input[0]
	isFirstCharSign := firstRune == '-' || firstRune == '+'
	processingPointer := input
	if isFirstCharSign {
		processingPointer = input[1:]
	}

	for _, r := range processingPointer {
		if !unicode.IsDigit(r) {
			return false
		}
	}

	return true
}

// Unsigned
//
// Example : https://play.golang.org/p/NohlCrIj77r
func (it *unicodeInteger) Unsigned(input string) bool {
	if input == "" || input == "+" || input == "-" || input[0] == '-' {
		return false
	}

	processingPointer := input
	firstChar := input[0]
	isPositiveSign := firstChar == '+'

	if isPositiveSign {
		processingPointer = input[1:]
	}

	for _, r := range processingPointer {
		if !unicode.IsDigit(r) {
			return false
		}
	}

	return true
}

// UnsignedWithoutPlus
//
// Returns true for positive number without `+` symbol as well.
func (it *unicodeInteger) UnsignedWithoutPlus(input string) bool {
	if input == "" || input == "+" || input == "-" || input[0] == '-' || input[0] == '+' {
		return false
	}

	for _, r := range input {
		if !unicode.IsDigit(r) {
			return false
		}
	}

	return true
}
