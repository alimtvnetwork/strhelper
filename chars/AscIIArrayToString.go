package chars

// Only return string for existing ones
func AscIIArrayToString(chars *[256]uint8) string {
	return string(*AscIIArrayToRunes(chars))
}
