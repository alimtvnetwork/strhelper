package strs

import (
	"bytes"
	"encoding/json"
)

// Returns:
//  - nil : if @anything is nil.
//  - *[]bytes : if anything exist and doesn't have any error from parsing json.NewEncoder(bytes.Buffer).Encode().
func ToBytesOfAny(anything interface{}) (*[]byte, error) {
	if anything == nil {
		return nil, nil
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
