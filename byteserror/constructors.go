package byteserror

import (
	"gitlab.com/evatix-go/strhelper/parsingtype"
	"gitlab.com/evatix-go/strhelper/strerror"
)

func NewError(err error, byteType parsingtype.ByteType) *Wrapper {
	return &Wrapper{
		errorWrapper: *strerror.NewErrorWrapperPtr(&err),
		byteType:     byteType,
	}
}

func NewErrorPtr(err *error, byteType parsingtype.ByteType) *Wrapper {
	return &Wrapper{
		errorWrapper: *strerror.NewErrorWrapperPtr(err),
		byteType:     byteType,
	}
}

func NewPtr(bytes *[]byte, err *error, byteType parsingtype.ByteType) *Wrapper {
	length := 0

	if bytes != nil {
		length = len(*bytes)
	}

	return &Wrapper{
		bytes:        bytes,
		errorWrapper: *strerror.NewErrorWrapperPtr(err),
		bytesLength:  length,
		byteType:     byteType,
	}
}

func New(bytes *[]byte, err error, byteType parsingtype.ByteType) Wrapper {
	length := 0

	if bytes != nil {
		length = len(*bytes)
	}

	return Wrapper{
		bytes:        bytes,
		errorWrapper: strerror.NewErrorWrapper(err),
		byteType:     byteType,
		bytesLength:  length,
	}
}

// NewNoError Creates new Wrapper
func NewNoError(bytes *[]byte, byteType parsingtype.ByteType) *Wrapper {
	return NewPtr(bytes, nil, byteType)
}

func EmptyPtr(byteType parsingtype.ByteType) *Wrapper {
	return &Wrapper{
		byteType: byteType,
	}
}

func Empty(byteType parsingtype.ByteType) Wrapper {
	return Wrapper{
		byteType: byteType,
	}
}
