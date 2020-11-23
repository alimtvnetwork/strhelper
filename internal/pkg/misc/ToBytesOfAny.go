package misc

import (
	"bytes"
	"encoding/json"
)

// Returns:
//  - nil : if @anything is nil.
//  - []bytes : if anything exist.  json.NewEncoder(bytes.Buffer).Encode....
func ToBytesOfAny(anything interface{}) *[]byte {
	if anything == nil {
		return nil
	}

	// Reference : https://stackoverflow.com/a/49946268
	reqBodyBytes := new(bytes.Buffer)
	encoder := json.NewEncoder(reqBodyBytes)
	encoder.Encode(anything)
	bytes := reqBodyBytes.Bytes()

	return &bytes
}
