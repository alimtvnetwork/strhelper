package strlines

import "unsafe"

func ToUnsafeBytes(lines *[]string) []byte {
	if lines == nil || *lines == nil {
		return nil
	}

	// https://stackoverflow.com/a/51953176
	rawBytesPtr := (*[]byte)(unsafe.Pointer(lines))

	if rawBytesPtr != nil {
		return *rawBytesPtr
	}

	return nil
}
