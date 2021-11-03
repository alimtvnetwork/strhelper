package isstrs

func EmptyPtr(lines *[]string) bool {
	return lines == nil || len(*lines) == 0
}
