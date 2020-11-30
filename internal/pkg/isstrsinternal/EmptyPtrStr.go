package isstrsinternal

func EmptyPtrStr(lines *[]*string) bool {
	return lines == nil || *lines == nil || len(*lines) == 0
}
