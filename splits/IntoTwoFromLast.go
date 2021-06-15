package splits

func IntoTwoFromLast(s, separator string, isCaseSensitive bool) (left, right string) {
	return IntoTwoFromLastPtr(&s, &separator, isCaseSensitive)
}
