package strto

func BytesArrayPtr(string string) *[]byte {
	val := []byte(string)

	return &val
}
