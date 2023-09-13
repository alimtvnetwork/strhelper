package anyto

import (
	"gitlab.com/auk-go/strhelper/byteserror"
	"gitlab.com/auk-go/strhelper/encodingbytetype"
)

func BytesWrapperOf(any interface{}, parsingType encodingbytetype.Variant) *byteserror.Wrapper {
	if any == nil {
		return byteserror.EmptyPtr(parsingType)
	}

	switch parsingType {
	case encodingbytetype.Unknown, encodingbytetype.Encoding:
		return BytesWrapper(any)
	case encodingbytetype.Unsafe:
		return UnSafeBytesWrapper(any)
	case encodingbytetype.AnyToValueStringBytes:
		return ValueBytesWrapper(any)
	case encodingbytetype.JsonParsing:
		return JsonBytesWrapper(any)
	default:
		panic(parsingBytesNotSupportMessage(parsingType))
	}
}
