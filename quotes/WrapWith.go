package quotes

// if isSkipOnExists true then if quotes are already exist then skip the wrapping process.
// Note : That it doesn't care about in middle quotes, it just wraps around.
func WrapWith(str string, quote Quote, isSkipOnExists bool) string {
	return WrapWithPtr(&str, quote, isSkipOnExists)
}
