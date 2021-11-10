package anyto

import (
	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/errorwrapper/errdata/errjson"
	"gitlab.com/evatix-go/strhelper/byteserror"
	"gitlab.com/evatix-go/strhelper/encodingbytetype"
)

func JsonBytesWrapper(any interface{}) *byteserror.Wrapper {
	jsonResult := corejson.NewFromAnyPtr(any)
	errJson := errjson.New.Result.Item(jsonResult)

	return byteserror.NewPtr(
		encodingbytetype.JsonParsing,
		jsonResult.Bytes,
		errJson.ErrorWrapper)
}
