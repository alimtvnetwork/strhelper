package anyto

import (
	"fmt"

	"gitlab.com/auk-go/core/constants"
)

func ValueBytesPtr(anything interface{}) *[]byte {
	if anything == nil {
		return nil
	}

	allBytes := []byte(fmt.Sprintf(constants.SprintValueFormat, anything))

	return &allBytes
}
