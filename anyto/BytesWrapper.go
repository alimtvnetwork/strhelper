package anyto

import (
	"gitlab.com/evatix-go/strhelper/byteserror"
	"gitlab.com/evatix-go/strhelper/parsingtype"
)

func BytesWrapper(anything interface{}) byteserror.Wrapper {
	allBytes, err := Bytes(anything)

	return byteserror.New(allBytes, err, parsingtype.Encoding)
}
