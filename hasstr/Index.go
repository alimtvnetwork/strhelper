package hasstr

// Index returns true if len(str)-1 >= index
func Index(str string, index int) bool {
	if str == "" {
		return false
	}

	return len(str)-1 >= index
}
