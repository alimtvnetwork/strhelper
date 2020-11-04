package quotes

import "gitlab.com/evatix-go/strhelper/constants"

func WhichQuote(char uint8, isLeft bool) QuoteStatus {
	otherQuoteStatus, has := otherQuoteCharsMaps[char]

	if !has {
		return EmptyQuoteStatus()
	}

	otherQuoteStatus.IsLeft = isLeft

	return otherQuoteStatus
}

var otherQuoteMaps = map[Quote]Quote{
	Single: Single,
	Double: Double,
}

var otherQuoteCharsMaps = map[uint8]QuoteStatus{
	constants.SingleQuoteSymbol: {
		IsQuoteFound: true,
		Found:        Single,
	},
	constants.DoubleQuoteSymbol: {
		IsQuoteFound: true,
		Found:        Double,
	},
}

func isEmptyStringPtr(str *string) bool {
	return str == nil || *str == constants.EmptyString || len(*str) == 0
}

// if isSkipOnExists true then if quotes are already exist then skip the wrapping process.
// Note : That it doesn't care about in middle quotes, it just wraps around.
func WrapWith(str string, quote Quote, isSkipOnExists bool) string {
	return WrapWithPtr(&str, quote, isSkipOnExists)
}

// if isSkipOnExists true then if quotes are already exist then skip the wrapping process.
// Note : That it doesn't care about in middle quotes, it just wraps around.
func WrapWithPtr(str *string, quote Quote, isSkipOnExists bool) string {
	if isEmptyStringPtr(str) || !isSkipOnExists {
		return quote.Wrap(str)
	}

	if HasBothWrappedWithPtr(str, quote) {
		// no need to modify
		return *str
	}

	singleQuoteStatus := getQuoteStatus(str)

	if singleQuoteStatus.IsQuoteFound && singleQuoteStatus.IsLeft {
		return *str + singleQuoteStatus.Found.String()
	}

	if singleQuoteStatus.IsQuoteFound && !singleQuoteStatus.IsLeft {
		return singleQuoteStatus.Found.String() + *str
	}

	return quote.Wrap(str)
}

func HasBothWrappedWithPtr(str *string, quote Quote) bool {
	if isEmptyStringPtr(str) || len(*str) <= 1 {
		return false
	}

	// has at least 2 chars
	length := len(*str)
	firstChar := (*str)[0]
	lastChar := (*str)[length-1]

	return quote.IsEqual(firstChar) && quote.IsEqual(lastChar)
}

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
