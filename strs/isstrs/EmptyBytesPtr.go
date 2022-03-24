package isstrs

func EmptyBytesPtr(lines *[]byte) bool {
	return lines == nil || len(*lines) == 0
}
