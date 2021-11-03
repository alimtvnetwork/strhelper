package quotes

// UnWrapWithPtr
//
// Note : It doesn't care about in quotes exist in middle of (str),
// it just unwrap from both sides if quotes are there.
func UnWrapWithPtr(str *string, quote Quote) string {
	if isEmptyStringPtr(str) {
		return *str
	}

	if HasBothWrappedWithPtr(str, quote) {
		return unWrapBoth(str)
	}

	singleQuoteStatus := getQuoteStatus(str)

	if singleQuoteStatus.IsQuoteFound {
		return unWrapSingle(str, singleQuoteStatus.IsLeft)
	}

	// return as is, there is nothing to modify
	return *str
}
