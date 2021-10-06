package anyto

import (
	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/errorwrapper/errdata/errjson"
	"gitlab.com/evatix-go/strhelper/byteserror"
	"gitlab.com/evatix-go/strhelper/encodingbytetype"
)

func JsonBytesWrapper(any interface{}) *byteserror.Wrapper {
	if any == nil {
		return byteserror.EmptyPtr(encodingbytetype.JsonParsing)
	}

	jsonResult := corejson.NewFromAny(any)
	errJson := errjson.New(jsonResult.Bytes, jsonResult.Error)

	return byteserror.NewPtr(encodingbytetype.JsonParsing, jsonResult.Bytes, errJson.ErrorWrapper)
}
