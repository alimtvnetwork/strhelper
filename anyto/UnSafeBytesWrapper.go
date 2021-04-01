package anyto

import (
	"gitlab.com/evatix-go/strhelper/byteserror"
	"gitlab.com/evatix-go/strhelper/encodingbytetype"
)

func UnSafeBytesWrapper(any interface{}) byteserror.Wrapper {
	if any == nil {
		return byteserror.Empty(encodingbytetype.Unsafe)
	}

	unsafeBytes := UnSafeBytes(any)

	if unsafeBytes == nil {
		return byteserror.Empty(encodingbytetype.Unsafe)
	}

	return *byteserror.NewNoError(unsafeBytes, encodingbytetype.Unsafe)
}
