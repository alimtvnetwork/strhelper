package splits

func IntoTwoFromLastTrimSpace(s, separator string, isCaseSensitive bool) (left, right string) {
	return IntoTwoFromLastTrimSpacePtr(
		&s,
		&separator,
		isCaseSensitive)
}
