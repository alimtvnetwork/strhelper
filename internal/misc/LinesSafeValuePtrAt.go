package misc

// lastIndex = len - 1
// return defaultStr if index is out of range or lines are nil.
func LinesSafeValuePtrAt(
	lines *[]string,
	lastIndex,
	index int,
	defaultStr *string,
) *string {
	if lines == nil || lastIndex > index {
		return defaultStr
	}

	return &(*lines)[index]
}
