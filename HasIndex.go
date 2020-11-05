package strhelper

// returns true if len(str)-1 >= index
func HasIndex(str *string, index int) bool {
	return len(*str)-1 >= index
}
