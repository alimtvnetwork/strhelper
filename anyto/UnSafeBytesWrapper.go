package anyto

import (
	"gitlab.com/evatix-go/strhelper/byteserror"
	"gitlab.com/evatix-go/strhelper/parsingtype"
)

func UnSafeBytesWrapper(any interface{}) byteserror.Wrapper {
	if any == nil {
		return byteserror.Empty(parsingtype.Unsafe)
	}

	unsafeBytes := UnSafeBytes(any)

	if unsafeBytes == nil {
		return byteserror.Empty(parsingtype.Unsafe)
	}

	return *byteserror.NewNoError(unsafeBytes, parsingtype.Unsafe)
}
