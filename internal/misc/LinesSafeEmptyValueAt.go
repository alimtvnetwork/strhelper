package misc

// lastIndex = len - 1
// return "" if index is out of range or lines are nil.
func LinesSafeEmptyValueAt(
	lines *[]string,
	lastIndex,
	index int,
) string {
	if lines == nil || lastIndex > index {
		return ""
	}

	return (*lines)[index]
}
