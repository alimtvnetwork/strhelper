package misc

import "unsafe"

// Returns:
//  - nil : if @lines are nil.
//  - []bytes : if anything exist. Usages unsafe pointer casting to get the bytes.
func ToBytes(lines *[]string) *[]byte {
	if lines == nil {
		return nil
	}

	return (*[]byte)(unsafe.Pointer(lines))
}
