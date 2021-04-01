package anyto

import (
	"encoding/json"

	"gitlab.com/evatix-go/strhelper/byteserror"
	"gitlab.com/evatix-go/strhelper/encodingbytetype"
)

func JsonBytesWrapper(any interface{}) byteserror.Wrapper {
	if any == nil {
		return byteserror.Empty(encodingbytetype.JsonParsing)
	}

	jsonBytes, err := json.Marshal(any)

	if jsonBytes != nil {
		return byteserror.New(&jsonBytes, err, encodingbytetype.JsonParsing)
	}

	return byteserror.New(nil, err, encodingbytetype.JsonParsing)
}
