package byteserror

import (
	"gitlab.com/evatix-go/core/issetter"
	"gitlab.com/evatix-go/errorwrapper"

	"gitlab.com/evatix-go/strhelper/parsingtype"
)

type WrapperDataModel struct {
	Bytes        *[]byte
	ErrorWrapper *errorwrapper.Wrapper
	ByteType     parsingtype.Variant
	BytesLength  int
	IsWhitespace issetter.Value
}
