package chars

// Return runes for only the existing ones.
func AsciiArrayToRunes(chars *[256]uint8) *[]rune {
	newChars := make([]rune, 0, len(*chars)/2)

	for index, existenceMark := range chars {
		if existenceMark == 1 {
			newChars = append(newChars, rune(index))
		}
	}

	return &newChars
}
