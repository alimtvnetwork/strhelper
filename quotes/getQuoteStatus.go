package quotes

func getQuoteStatus(str *string) QuoteStatus {
	if isEmptyStringPtr(str) {
		return EmptyQuoteStatus()
	}

	// has at least 2 chars
	length := len(*str)
	firstChar := (*str)[0]
	whichQuoteFirstChar := WhichQuote(firstChar, true)

	if length == 1 || whichQuoteFirstChar.IsQuoteFound {
		return whichQuoteFirstChar
	}

	lastChar := (*str)[length-1]

	return WhichQuote(lastChar, false)
}
