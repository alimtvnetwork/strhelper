package strto

func RunesArrayPtr(str string) *[]rune {
	val := []rune(str)

	return &val
}
