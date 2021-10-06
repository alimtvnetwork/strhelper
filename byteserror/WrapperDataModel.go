package byteserror

import (
	"gitlab.com/evatix-go/core/issetter"
	"gitlab.com/evatix-go/errorwrapper"

	"gitlab.com/evatix-go/strhelper/encodingbytetype"
)

type WrapperDataModel struct {
	Bytes        []byte
	ErrorWrapper *errorwrapper.Wrapper
	ByteType     encodingbytetype.Variant
	BytesLength  int
	IsWhitespace issetter.Value
}
