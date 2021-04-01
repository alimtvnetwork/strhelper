package byteserror

import (
	"gitlab.com/evatix-go/errorwrapper/errnew"

	"gitlab.com/evatix-go/strhelper/encodingbytetype"
)

func NewError(err error, byteType encodingbytetype.Variant) *Wrapper {
	return &Wrapper{
		errorWrapper: errnew.ErrPtr(err),
		byteType:     byteType,
	}
}

func NewErrorPtr(err *error, byteType encodingbytetype.Variant) *Wrapper {
	return &Wrapper{
		errorWrapper: errnew.ErrInPtr(err),
		byteType:     byteType,
	}
}

func NewPtr(bytes *[]byte, err *error, byteType encodingbytetype.Variant) *Wrapper {
	length := 0

	if bytes != nil {
		length = len(*bytes)
	}

	return &Wrapper{
		bytes:        bytes,
		errorWrapper: errnew.ErrInPtr(err),
		bytesLength:  length,
		byteType:     byteType,
	}
}

func New(bytes *[]byte, err error, byteType encodingbytetype.Variant) Wrapper {
	length := 0

	if bytes != nil {
		length = len(*bytes)
	}

	return Wrapper{
		bytes:        bytes,
		errorWrapper: errnew.ErrPtr(err),
		byteType:     byteType,
		bytesLength:  length,
	}
}

// NewNoError Creates new Wrapper
func NewNoError(bytes *[]byte, byteType encodingbytetype.Variant) *Wrapper {
	return NewPtr(bytes, nil, byteType)
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
