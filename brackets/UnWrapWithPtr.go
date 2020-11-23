package brackets

// Note : It doesn't care about in brackets exist in middle of (str),
// it just unwrap from both sides if brackets are there.
func UnWrapWithPtr(str *string, category Category) string {
	if isEmptyStringPtr(str) {
		return *str
	}

	if HasBothWrappedWithPtr(str, category) {
		// no need to modify
		return unWrapBoth(str)
	}

	status := getSingleBracketStatus(str)

	if status.IsBracketFound {
		return unWrapSingle(str, status.IsLeft)
	}

	// return as is, there is nothing to modify
	return *str
}
