package chars

// Modifies existing chars array to reverse order.
// if nil or empty then returns nil
func ReverseInPlacePtr(chars *[]uint8) *[]uint8 {
	length := len(*chars)

	if length == 0 {
		return nil
	}

	mid := length / 2
	lastIndex := length - 1

	for i := 0; i < mid; i++ {
		(*chars)[i], (*chars)[lastIndex-i] = (*chars)[lastIndex-i], (*chars)[i]
	}

	return chars
}
