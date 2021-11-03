package strlines

// lastIndex = len - 1
// return defaultStr if index is out of range or strlines are nil.
func SafeValueAt(
	lines []string,
	lastIndex,
	index int,
	defaultStr string,
) string {
	if lines == nil || lastIndex > index {
		return defaultStr
	}

	return lines[index]
}
