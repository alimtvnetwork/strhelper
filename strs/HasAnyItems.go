package strs

// HasAnyItems
//
// Refers to non-empty array !(lines == nil || len(*lines) == 0)
func HasAnyItems(lines []string) bool {
	return len(lines) > 0
}
