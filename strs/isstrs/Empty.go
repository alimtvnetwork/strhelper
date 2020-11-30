package isstrs

func Empty(lines *[]string) bool {
	return lines == nil || *lines == nil || len(*lines) == 0
}
