package anyto

import (
	"fmt"

	"gitlab.com/evatix-go/strhelper/byteserror"
	"gitlab.com/evatix-go/strhelper/parsingtype"
)

func ValueBytesWrapper(anything interface{}) *byteserror.Wrapper {
	if anything == nil {
		return nil
	}

	allBytes := []byte(fmt.Sprintf("%v", anything))

	return byteserror.NewNoError(&allBytes, parsingtype.AnyToValueStringBytes)
}
