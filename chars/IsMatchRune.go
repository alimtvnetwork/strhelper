package chars

func IsMatchRune(rune1 rune, rune2 rune, isCaseSensitive bool) bool {
	if isCaseSensitive {
		return rune1 == rune2
	}

	// Insensitive case
	return ToLowerRune(rune1) == ToLowerRune(rune2)
}
