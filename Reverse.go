package strhelper

import "gitlab.com/evatix-go/strhelper/constants"

// Returns empty string if str is nil or empty.
// Returns new string which is reversed.
func Reverse(str string) string {
	return ReversePtr(&str)
}

// Returns empty string if str is nil or empty.
// Returns new string which is reversed.
func ReversePtr(str *string) string {
	length := len(*str)

	if length == 0 {
		return constants.EmptyString
	}

	runes := []rune(*str)

	return string(*ReverseRuneInPlacePtr(&runes))
}

// Modifies existing runesIn array to reverse order.
// if nil or empty then return nil
func ReverseRuneInPlacePtr(runesIn *[]rune) *[]rune {
	length := len(*runesIn)

	if length == 0 {
		return nil
	}

	mid := length / 2
	lastIndex := length - 1

	for i := 0; i < mid; i++ {
		(*runesIn)[i], (*runesIn)[lastIndex-i] = (*runesIn)[lastIndex-i], (*runesIn)[i]
	}

	return runesIn
}
