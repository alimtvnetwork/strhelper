package concat

func StringsWithSeparator(
	currentStr,
	separator string,
	isSkipEmptyOrNil bool,
	contents ...string,
) string {
	return StringsArrayWithSeparator(
		&currentStr,
		&separator,
		isSkipEmptyOrNil,
		&contents)
}
