package strlines

// lastIndex = len - 1
// return "" if index is out of range or strlines are nil.
func SafeEmptyValueAt(
	lines []string,
	lastIndex,
	index int,
) string {
	if lines == nil || lastIndex > index {
		return ""
	}

	return lines[index]
}
