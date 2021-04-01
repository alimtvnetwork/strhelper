package strvalidator

func HasNegativeStart(str string) bool {
	return !(str == "" || str[0] != '-')
}
