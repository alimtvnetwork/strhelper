package brackets

// if isSkipOnExists true then if both brackets are exist then skip the process or
// if single is there then find the next one and apply (may not apply properly if missing)
func WrapWithPtr(str *string, category Category, isSkipOnExists bool) string {
	if !isSkipOnExists || isEmptyStringPtr(str) {
		return category.Pair().Wrap(str)
	}

	if HasBothWrappedWithPtr(str, category) && isSkipOnExists {
		// no need to modify
		return *str
	}

	singleBracketStatus := getSingleBracketStatus(str)

	if singleBracketStatus.IsBracketFound && singleBracketStatus.IsLeft {
		// no need to modify
		return *str + singleBracketStatus.OtherBracket.String()
	}

	if singleBracketStatus.IsBracketFound && !singleBracketStatus.IsLeft {
		// no need to modify
		return singleBracketStatus.OtherBracket.String() + *str
	}

	return category.Pair().Wrap(str)
}
