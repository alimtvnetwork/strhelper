package chars

// Makes a new char to Upper, doesn't modify the existing one.
//
// Warning: Requires double memory cost for copying the same data.
func ToCharsUpper(chars *[]uint8) *[]uint8 {
	// Copying, reference: https://play.golang.org/p/r65MrCg86YH
	newChars := *chars

	return ToCharsUpperInPlace(&newChars)
}
