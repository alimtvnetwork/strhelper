package chars

func IsMatchRuneCaseInsensitive(rune1 rune, rune2 rune) bool {
	return ToLowerRune(rune1) == ToLowerRune(rune2)
}
