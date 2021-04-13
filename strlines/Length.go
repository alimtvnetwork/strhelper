package strlines

// For nil return 0 no panic.
func Length(lines *[]string) int {
	if lines == nil {
		return 0
	}

	return len(*lines)
}
