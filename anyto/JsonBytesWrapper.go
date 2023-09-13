package anyto

import (
	"gitlab.com/auk-go/core/coredata/corejson"
	"gitlab.com/auk-go/errorwrapper/errdata/errjson"
	"gitlab.com/auk-go/strhelper/byteserror"
	"gitlab.com/auk-go/strhelper/encodingbytetype"
)

func JsonBytesWrapper(any interface{}) *byteserror.Wrapper {
	jsonResult := corejson.NewPtr(any)
	errJson := errjson.New.Result.Item(jsonResult)

	return byteserror.NewPtr(
		encodingbytetype.JsonParsing,
		jsonResult.Bytes,
		errJson.ErrorWrapper)
}
