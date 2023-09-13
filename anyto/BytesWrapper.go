package anyto

import (
	"gitlab.com/auk-go/core/codestack"
	"gitlab.com/auk-go/errorwrapper/errnew"
	"gitlab.com/auk-go/errorwrapper/errtype"
	"gitlab.com/auk-go/strhelper/byteserror"
	"gitlab.com/auk-go/strhelper/encodingbytetype"
)

func BytesWrapper(anything interface{}) *byteserror.Wrapper {
	allBytes, err := Bytes(anything)

	errWp := errnew.Error.UsingStackSkip(
		codestack.Skip1,
		errtype.ConversionFailed,
		err)

	return byteserror.NewPtr(
		encodingbytetype.Encoding,
		allBytes,
		errWp)
}
