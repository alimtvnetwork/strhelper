package chars

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
