package charhelper

// Makes a new chars to reverse, doesn't modify the existing one.
// if nil or empty then return nil
func ReversePtr(chars *[]uint8) *[]uint8 {
	length := len(*chars)

	if length == 0 {
		return nil
	}

	newChars := make([]uint8, length)
	mid := length / 2
	lastIndex := length - 1

	for i := 0; i < mid; i++ {
		newChars[i], newChars[lastIndex-i] = (*chars)[lastIndex-i], (*chars)[i]
	}

	if length%2 == 0 {
		newChars[mid] = (*chars)[mid]
	}

	return &newChars
}
