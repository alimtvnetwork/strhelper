package reverse

import (
	"gitlab.com/evatix-go/strhelper/strconst"
)

// Returns empty string if str is nil or empty.
// Returns new string which is reversed.
func Get(str string) string {
	return Ptr(&str)
}

// Returns empty string if str is nil or empty.
// Returns new string which is reversed.
func Ptr(str *string) string {
	length := len(*str)

	if length == 0 {
		return strconst.EmptyString
	}

	runes := []rune(*str)

	return string(*RunesInPlacePtr(&runes))
}

// Modifies existing runesIn array to reverse order.
// if nil or empty then returns nil
func RunesInPlacePtr(runesIn *[]rune) *[]rune {
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
