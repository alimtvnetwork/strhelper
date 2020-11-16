package chars

// Makes a new chars to reverse, doesn't modify the existing one.
// if nil or empty then return nil
func Reverse(chars []uint8) []uint8 {
	return *ReversePtr(&chars)
}
