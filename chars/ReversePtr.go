package chars

// Makes a new chars to reverse, doesn't modify the existing one.
// if nil or empty then return nil
func ReversePtr(chars *[]uint8) *[]uint8 {
	length := len(*chars)

	if length == 0 {
		return nil
	}

	// Copying, reference: https://play.golang.org/p/r65MrCg86YH
	newChars := *chars

	return ReverseInPlacePtr(&newChars)
}
