package chars

func ToRuneArrayPtr(string *string) *[]rune {
	val := []rune(*string)

	return &val
}
