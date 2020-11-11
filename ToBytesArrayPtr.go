package strhelper

func ToBytesArrayPtr(string string) *[]byte {
	val := []byte(string)

	return &val
}
