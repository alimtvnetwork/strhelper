package anyto

import (
	"bytes"
	"encoding/json"

	"gitlab.com/evatix-go/strhelper/internal/reflectinternal"
)

// Returns:
//  - nil : if @anything is nil.
//  - *[]bytes : if anything exist and doesn't have any error from parsing json.NewEncoder(bytes.Buffer).Encode().
func Bytes(anything interface{}) (*[]byte, error) {
	if anything == nil {
		return nil, nil
	}

	isBytes, allBytes := reflectinternal.IsBytesOrBytesPointer(anything)

	if isBytes {
		return allBytes, nil
	}

	isString, str := reflectinternal.IsString(anything)

	if isString {
		toBytes := []byte(str)

		return &toBytes, nil
	}

	isByte, currentByte := reflectinternal.IsByte(anything)

	if isByte {
		toBytes := []byte{currentByte}

		return &toBytes, nil
	}

	// Reference : https://stackoverflow.com/a/49946268
	reqBodyBytes := new(bytes.Buffer)
	encoder := json.NewEncoder(reqBodyBytes)
	err := encoder.Encode(anything)

	if err != nil {
		return nil, err
	}

	currentBytes := reqBodyBytes.Bytes()

	if currentBytes != nil {
		return &currentBytes, nil
	}

	return nil, nil
}
