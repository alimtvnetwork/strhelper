package quotes

// if isSkipOnExists true then if quotes are already exist then skip the wrapping process.
// Note : That it doesn't care about in middle quotes, it just wraps around.
func WrapWithPtr(str *string, quote Quote, isSkipOnExists bool) string {
	if isEmptyStringPtr(str) {
		return quote.Wrap(str)
	}

	if HasBothWrappedWithPtr(str, quote) && isSkipOnExists {
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
