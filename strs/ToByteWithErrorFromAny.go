package strs

import (
	"bytes"
	"encoding/json"
	"errors"

	"gitlab.com/evatix-go/strhelper/strhelpercore"
)

// Returns:
//  - nil : if @anything is nil.
//  - strhelpercore.BytesWithError : if anything exist and conversion
//      has no error from json.NewEncoder(bytes.Buffer).Encode
func ToByteWithErrorFromAny(anything interface{}) *strhelpercore.BytesWithError {
	if anything == nil {
		return strhelpercore.NewBytesWithErrorOnlyError(errors.New("ToByteWithErrorFromAny conversion object given as nil"))
	}

	// Reference : https://stackoverflow.com/a/49946268
	reqBodyBytes := new(bytes.Buffer)
	encoder := json.NewEncoder(reqBodyBytes)
	err := encoder.Encode(anything)

	if err != nil {
		return strhelpercore.NewBytesWithErrorOnlyError(err)
	}

	currentBytes := reqBodyBytes.Bytes()

	if currentBytes != nil {
		return strhelpercore.NewBytesWithNoError(currentBytes)
	}

	return strhelpercore.NewBytesWithNoError(nil)
}
