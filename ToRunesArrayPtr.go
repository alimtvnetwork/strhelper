package strhelper

func ToRunesArrayPtr(string string) *[]rune {
	val := []rune(string)

	return &val
}
