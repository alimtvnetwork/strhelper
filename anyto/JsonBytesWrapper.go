package anyto

import (
	"encoding/json"

	"gitlab.com/evatix-go/strhelper/byteserror"
	"gitlab.com/evatix-go/strhelper/parsingtype"
)

func JsonBytesWrapper(any interface{}) byteserror.Wrapper {
	if any == nil {
		return byteserror.Empty(parsingtype.JsonParsing)
	}

	jsonBytes, err := json.Marshal(any)

	if jsonBytes != nil {
		return byteserror.New(&jsonBytes, err, parsingtype.JsonParsing)
	}

	return byteserror.New(nil, err, parsingtype.JsonParsing)
}
