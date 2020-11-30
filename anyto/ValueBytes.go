package anyto

import "fmt"

func ValueBytes(anything interface{}) []byte {
	if anything == nil {
		return nil
	}

	return []byte(fmt.Sprintf("%v", anything))
}
