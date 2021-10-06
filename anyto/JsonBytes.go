package anyto

import "encoding/json"

func JsonBytes(any interface{}) []byte {
	if any == nil {
		return nil
	}

	jsonBytes, err := json.Marshal(any)

	if err != nil {
		return []byte{}
	}

	return jsonBytes
}
