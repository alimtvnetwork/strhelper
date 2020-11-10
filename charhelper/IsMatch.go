package charhelper

func IsMatch(char1 uint8, char2 uint8, isCaseSensitive bool) bool {
	if isCaseSensitive {
		return char1 == char2
	}

	// Insensitive case
	return ToLower(char1) == ToLower(char2)
}
