package hasstr

// returns true if len(str) >= length
func Length(str *string, length int) bool {
	return len(*str) >= length
}
