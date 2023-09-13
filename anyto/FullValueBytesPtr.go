package anyto

import (
	"fmt"

	"gitlab.com/auk-go/core/constants"
)

func FullValueBytesPtr(anything interface{}) *[]byte {
	if anything == nil {
		return nil
	}

	allBytes := []byte(
		fmt.Sprintf(constants.SprintFullPropertyNameValueFormat,
			anything))

	return &allBytes
}
