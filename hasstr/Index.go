package hasstr

// returns true if len(str)-1 >= index
func Index(str *string, index int) bool {
	return len(*str)-1 >= index
}
