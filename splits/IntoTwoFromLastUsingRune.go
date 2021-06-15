package splits

func IntoTwoFromLastUsingRune(s string, splitRune rune) (left, right string) {
	return IntoTwoFromLastUsingRunePtr(
		&s,
		splitRune)
}
