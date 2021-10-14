package anyto

import (
	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/core/errcore"
	"gitlab.com/evatix-go/strhelper/encodingbytetype"
)

func BytesOf(parsingType encodingbytetype.Variant, any interface{}) ([]byte, error) {
	if any == nil {
		return nil, nil
	}

	switch parsingType {
	case encodingbytetype.Unknown, encodingbytetype.Encoding:
		return Bytes(any)
	case encodingbytetype.Unsafe:
		return UnSafeBytes(any), nil
	case encodingbytetype.AnyToValueStringBytes:
		return ValueBytes(any), nil
	case encodingbytetype.AnyToFullStringBytes:
		return FullValueBytes(any), nil
	case encodingbytetype.JsonParsing:
		jsonResult := corejson.NewFromAny(any)

		return jsonResult.Bytes, jsonResult.MeaningfulError()
	default:
		return nil, errcore.NotSupportedType.Error(
			parsingBytesNotSupportMessage(parsingType), any)
	}
}

func parsingBytesNotSupportMessage(parsingType encodingbytetype.Variant) string {
	return "Parsing type not support for bytes conversion. Requested parsing type : " + parsingType.String()
}
