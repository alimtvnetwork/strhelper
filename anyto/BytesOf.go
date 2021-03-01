package anyto

import (
	"gitlab.com/evatix-go/strhelper/parsingtype"
)

func BytesOf(any interface{}, parsingType parsingtype.Variant) (*[]byte, error) {
	if any == nil {
		return nil, nil
	}

	switch parsingType {
	case parsingtype.Unknown, parsingtype.Encoding:
		return Bytes(any)
	case parsingtype.Unsafe:
		return UnSafeBytes(any), nil
	case parsingtype.AnyToValueStringBytes:
		return ValueBytesPtr(any), nil
	case parsingtype.AnyToFullStringBytes:
		return FullValueBytesPtr(any), nil
	case parsingtype.JsonParsing:
		return JsonBytes(any), nil
	default:
		panic(parsingBytesNotSupportMessage(parsingType))
	}
}

func parsingBytesNotSupportMessage(parsingType parsingtype.Variant) string {
	return "Parsing type not support for bytes conversion. Requested parsing type : " + parsingType.String()
}
