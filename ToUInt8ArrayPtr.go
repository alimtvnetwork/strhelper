package strhelper

func ToUInt8ArrayPtr(string string) *[]uint8 {
	val := []uint8(string)

	return &val
}
