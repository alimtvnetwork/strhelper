package strs

func IsEmpty(lines *[]string) bool {
	return lines == nil || *lines == nil || len(*lines) == 0
}
