package anyto

import (
	"fmt"

	"gitlab.com/auk-go/strhelper/byteserror"
	"gitlab.com/auk-go/strhelper/encodingbytetype"
)

func ValueBytesWrapper(anything interface{}) *byteserror.Wrapper {
	if anything == nil {
		return nil
	}

	allBytes := []byte(fmt.Sprintf("%v", anything))

	return byteserror.NewNoErrorPtr(
		encodingbytetype.AnyToValueStringBytes,
		allBytes)
}
