package chars

func StringFromRunes(runes *[]rune) *string {
	value := string(*runes)

	return &value
}
