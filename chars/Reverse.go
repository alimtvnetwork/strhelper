package chars

// Reverse
//
// Makes a new chars to reverse, doesn't modify the existing one.
// if nil or empty then returns nil
func Reverse(chars []uint8) []uint8 {
	length := len(chars)

	if length == 0 {
		return nil
	}

	mid := length / 2
	lastIndex := length - 1
	toFinalList := make([]uint8, length)

	for i := 0; i < mid; i++ {
		toFinalList[i], toFinalList[lastIndex-i] = chars[lastIndex-i], chars[i]
	}

	return toFinalList
}
