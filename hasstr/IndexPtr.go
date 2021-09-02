package hasstr

// IndexPtr returns true if len(str)-1 >= index
func IndexPtr(str *string, index int) bool {
	if str == nil {
		return false
	}

	return len(*str)-1 >= index
}
