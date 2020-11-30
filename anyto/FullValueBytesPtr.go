package anyto

import (
	"fmt"

	"gitlab.com/evatix-go/strhelper/strconst"
)

func FullValueBytesPtr(anything interface{}) *[]byte {
	if anything == nil {
		return nil
	}

	allBytes := []byte(fmt.Sprintf(strconst.SprintFullPropertyNameValueFormat, anything))

	return &allBytes
}
