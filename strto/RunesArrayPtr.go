package strto

func RunesArrayPtr(string string) *[]rune {
	val := []rune(string)

	return &val
}
