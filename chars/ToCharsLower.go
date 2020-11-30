package chars

// Makes a new char to lower, doesn't modify the existing one.
func ToCharsLower(chars *[]uint8) *[]uint8 {
	// Copying, example reference : https://play.golang.org/p/r65MrCg86YH
	newChars := *chars

	return ToCharsLowerInPlace(&newChars)
}
