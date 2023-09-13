package byteserror

import (
	"gitlab.com/auk-go/core/issetter"
	"gitlab.com/auk-go/errorwrapper"

	"gitlab.com/auk-go/strhelper/encodingbytetype"
)

type WrapperDataModel struct {
	Bytes        []byte
	ErrorWrapper *errorwrapper.Wrapper
	ByteType     encodingbytetype.Variant
	BytesLength  int
	IsWhitespace issetter.Value
}
