package strvalidator

func HasNegativeStart(str *string) bool {
	return !(str == nil || *str == "" || (*str)[0] != '-')
}
