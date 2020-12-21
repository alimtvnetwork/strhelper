package misc

// For nil return 0 no panic.
func LinesLength(lines *[]string) int {
	if lines == nil {
		return 0
	}

	return len(*lines)
}
