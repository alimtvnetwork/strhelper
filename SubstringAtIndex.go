package strhelper

// language integrated ones will be faster str[startAtIndex:endsAtIndex]
// Under the hood this method usages that functionality from language
// panics if startsAtIndex < 0
func SubstringAtIndex(
	str string,
	startsAtIndex, endsAtIndex int,
) string {
	if startsAtIndex < 0 {
		message := "Substring Index cannot have negative startsAtIndex : " + IntToString(startsAtIndex)

		panic(message)
	}

	return str[startsAtIndex:endsAtIndex]
}
