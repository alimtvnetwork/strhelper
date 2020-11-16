package strto

func UInt8ArrayPtr(string string) *[]uint8 {
	val := []uint8(string)

	return &val
}
