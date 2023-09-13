package byteserror

import (
	"gitlab.com/auk-go/core/codestack"
	"gitlab.com/auk-go/errorwrapper"
	"gitlab.com/auk-go/errorwrapper/errnew"
	"gitlab.com/auk-go/errorwrapper/errtype"

	"gitlab.com/auk-go/strhelper/encodingbytetype"
)

func NewErrorPtr(
	byteType encodingbytetype.Variant,
	errType errtype.Variation,
	err error,
) *Wrapper {
	return &Wrapper{
		errorWrapper: errnew.Error.UsingStackSkip(
			codestack.Skip1,
			errType,
			err),
		byteType: byteType,
	}
}

func NewPtr(
	byteType encodingbytetype.Variant,
	bytes []byte,
	errWrapper *errorwrapper.Wrapper,
) *Wrapper {
	length := len(bytes)

	return &Wrapper{
		bytes:        bytes,
		errorWrapper: errWrapper,
		bytesLength:  length,
		byteType:     byteType,
	}
}

// NewNoErrorPtr
//
// NewNoError Creates new Wrapper
func NewNoErrorPtr(byteType encodingbytetype.Variant, allBytes []byte) *Wrapper {
	return NewPtr(byteType, allBytes, nil)
}

func EmptyPtr(byteType encodingbytetype.Variant) *Wrapper {
	return &Wrapper{
		byteType: byteType,
	}
}

func Empty(byteType encodingbytetype.Variant) Wrapper {
	return Wrapper{
		byteType: byteType,
	}
}
