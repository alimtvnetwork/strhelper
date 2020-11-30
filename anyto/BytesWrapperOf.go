package anyto

import (
	"gitlab.com/evatix-go/strhelper/byteserror"
	"gitlab.com/evatix-go/strhelper/parsingtype"
)

func BytesWrapperOf(any interface{}, parsingType parsingtype.ByteType) byteserror.Wrapper {
	if any == nil {
		return byteserror.Empty(parsingType)
	}

	switch parsingType {
	case parsingtype.Unknown, parsingtype.Encoding:
		return BytesWrapper(any)
	case parsingtype.Unsafe:
		return UnSafeBytesWrapper(any)
	case parsingtype.AnyToValueStringBytes:
		return *ValueBytesWrapper(any)
	case parsingtype.JsonParsing:
		return JsonBytesWrapper(any)
	default:
		panic(parsingBytesNotSupportMessage(parsingType))
	}
}
