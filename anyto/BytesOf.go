package anyto

import (
	"gitlab.com/evatix-go/strhelper/encodingbytetype"
)

func BytesOf(any interface{}, parsingType encodingbytetype.Variant) (*[]byte, error) {
	if any == nil {
		return nil, nil
	}

	switch parsingType {
	case encodingbytetype.Unknown, encodingbytetype.Encoding:
		return Bytes(any)
	case encodingbytetype.Unsafe:
		return UnSafeBytes(any), nil
	case encodingbytetype.AnyToValueStringBytes:
		return ValueBytesPtr(any), nil
	case encodingbytetype.AnyToFullStringBytes:
		return FullValueBytesPtr(any), nil
	case encodingbytetype.JsonParsing:
		return JsonBytes(any), nil
	default:
		panic(parsingBytesNotSupportMessage(parsingType))
	}
}

func parsingBytesNotSupportMessage(parsingType encodingbytetype.Variant) string {
	return "Parsing type not support for bytes conversion. Requested parsing type : " + parsingType.String()
}
