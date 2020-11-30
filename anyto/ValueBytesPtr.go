package anyto

import (
	"fmt"

	"gitlab.com/evatix-go/strhelper/strconst"
)

func ValueBytesPtr(anything interface{}) *[]byte {
	if anything == nil {
		return nil
	}

	allBytes := []byte(fmt.Sprintf(strconst.SprintValueFormat, anything))

	return &allBytes
}
