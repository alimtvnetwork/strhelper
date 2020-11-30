package anyto

import "encoding/json"

func JsonBytes(any interface{}) *[]byte {
	if any == nil {
		return nil
	}

	jsonBytes, err := json.Marshal(any)

	if err != nil {
		panic(err)
	}

	if jsonBytes != nil {
		return &jsonBytes
	}

	return nil
}
