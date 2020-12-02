package isstrsinternal

func EmptyPtr(lines *[]string) bool {
	return lines == nil || *lines == nil || len(*lines) == 0
}
