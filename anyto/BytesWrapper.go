package anyto

import (
	"gitlab.com/evatix-go/core/codestack"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/strhelper/byteserror"
	"gitlab.com/evatix-go/strhelper/encodingbytetype"
)

func BytesWrapper(anything interface{}) *byteserror.Wrapper {
	allBytes, err := Bytes(anything)

	errWp := errnew.UsingStackSkip(
		codestack.Skip1,
		errtype.ConversionFailed,
		err)

	return byteserror.NewPtr(
		encodingbytetype.Encoding,
		allBytes,
		errWp)
}
