package strs

func IsEmptyPtrStr(lines *[]*string) bool {
	return lines == nil || *lines == nil || len(*lines) == 0
}
