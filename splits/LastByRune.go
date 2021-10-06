package splits

func LastByRune(
	s string,
	runeToSplit rune,
	limits int,
) []string {
	return LastByRunePtr(
		s,
		runeToSplit,
		limits)
}
