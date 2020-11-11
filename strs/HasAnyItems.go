package strs

// Refers to non empty array !(lines == nil || len(*lines) == 0)
func HasAnyItems(lines *[]string) bool {
	return !(lines == nil || len(*lines) == 0)
}
