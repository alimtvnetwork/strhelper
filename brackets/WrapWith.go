package brackets

// if isSkipOnExists true then if both brackets are exist then skip the process or
// if single is there then find the next one and apply (may not apply properly if missing)
func WrapWith(str string, category Category, isSkipOnExists bool) string {
	return WrapWithPtr(&str, category, isSkipOnExists)
}
