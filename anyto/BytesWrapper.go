package anyto

import (
	"gitlab.com/evatix-go/strhelper/byteserror"
	"gitlab.com/evatix-go/strhelper/encodingbytetype"
)

func BytesWrapper(anything interface{}) byteserror.Wrapper {
	allBytes, err := Bytes(anything)

	return byteserror.New(allBytes, err, encodingbytetype.Encoding)
}
