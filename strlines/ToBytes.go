package strlines

import "unsafe"

// Returns:
//  - nil : if @strlines are nil.
//  - []bytes : if anything exist. Usages unsafe pointer casting to get the bytes.
// @Example:
//  - https://play.golang.org/p/cGpfXgZ3ZeC
// Expression:
//  - return (*[]byte)(unsafe.Pointer(strlines))
func ToUnsafeBytes(lines *[]string) *[]byte {
	if lines == nil || *lines == nil {
		return nil
	}

	// https://stackoverflow.com/a/51953176
	return (*[]byte)(unsafe.Pointer(lines))
}
