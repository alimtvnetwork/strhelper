package anyto

func UnSafeBytes(any interface{}) []byte {
	rawBytes := UnSafeBytesPtr(any)

	if rawBytes == nil || *rawBytes == nil {
		return []byte{}
	}

	return *rawBytes
}
