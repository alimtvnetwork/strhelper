package chars

// AsciiArrayToString Only return string for existing ones
func AsciiArrayToString(chars *[256]uint8) string {
	return string(*AsciiArrayToRunes(chars))
}
