package strhelper

// returns true if len(str) >= length
func HasLength(str *string, length int) bool {
	return len(*str) >= length
}
