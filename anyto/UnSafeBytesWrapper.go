package anyto

import (
	"gitlab.com/evatix-go/strhelper/byteserror"
	"gitlab.com/evatix-go/strhelper/encodingbytetype"
)

func UnSafeBytesWrapper(any interface{}) *byteserror.Wrapper {
	if any == nil {
		return byteserror.EmptyPtr(encodingbytetype.Unsafe)
	}

	unsafeBytes := UnSafeBytes(any)

	if unsafeBytes == nil {
		return byteserror.EmptyPtr(encodingbytetype.Unsafe)
	}

	return byteserror.NewNoErrorPtr(
		encodingbytetype.Unsafe,
		unsafeBytes)
}
