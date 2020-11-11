package brackets

func HasBothWrappedWithPtr(str *string, category Category) bool {
	if isEmptyStringPtr(str) || len(*str) <= 1 {
		return false
	}

	// has at least 2 chars
	length := len(*str)
	pair := category.Pair()
	firstChar := (*str)[0]
	lastChar := (*str)[length-1]

	return pair.Start.IsEqual(firstChar) && pair.End.IsEqual(lastChar)
}
