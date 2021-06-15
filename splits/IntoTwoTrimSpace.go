package splits

func IntoTwoTrimSpace(s, separator string) (left, right string) {
	return IntoTwoTrimSpacePtr(&s, &separator)
}
