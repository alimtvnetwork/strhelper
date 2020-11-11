package brackets

// This method is called when both brackets are not found.
// So the assumptions are first both bracket is done.
// Now possibility here, one exist another not
func getSingleBracketStatus(str *string) BracketStatus {
	if isEmptyStringPtr(str) {
		return EmptyBracketStatus()
	}

	// has at least 2 chars
	length := len(*str)
	firstChar := (*str)[0]
	whichBracketFirstChar := WhichBracket(firstChar, true)

	if length == 1 || whichBracketFirstChar.IsBracketFound {
		return whichBracketFirstChar
	}

	lastChar := (*str)[length-1]

	return WhichBracket(lastChar, false)
}
