package misc

func EmptyBytes(lines *[]byte) bool {
	return lines == nil || *lines == nil || len(*lines) == 0
}
