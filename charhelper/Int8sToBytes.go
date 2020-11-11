package charhelper

// Reference : https://stackoverflow.com/a/28848879
func Int8sToBytes(int8s *[]int8) []byte {
	b := make([]byte, len(*int8s))
	for i, v := range *int8s {
		b[i] = byte(v)
	}

	return b
}
