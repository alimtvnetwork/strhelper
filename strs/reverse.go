package strs

// Returns len == 0 then returns as is.
// Returns new strings which is reversed.
func Reverse(inputStrings []string) *[]string {
	return ReverseStringsPtr(&inputStrings)
}

// Returns len == 0 then returns as is.
// Returns new strings which is reversed from the given array.
// Reference array copy/referencing : https://play.golang.org/p/GwELxdHZWoy
func ReverseStringsPtr(inputStrings *[]string) *[]string {
	length := len(*inputStrings)

	if length == 0 {
		return inputStrings
	}

	// an example : https://play.golang.org/p/GwELxdHZWoy
	copyReference := *inputStrings

	return ReverseStringsInPlacePtr(&copyReference)
}

// Returns len == 0 then returns as is.
// Modifies existing strings array to reverse order.
func ReverseStringsInPlacePtr(inputStrings *[]string) *[]string {
	length := len(*inputStrings)

	if length == 0 {
		return nil
	}

	mid := length / 2
	lastIndex := length - 1

	for i := 0; i < mid; i++ {
		(*inputStrings)[i], (*inputStrings)[lastIndex-i] = (*inputStrings)[lastIndex-i], (*inputStrings)[i]
	}

	return inputStrings
}
