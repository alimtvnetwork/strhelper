package strhelper

import (
	"fmt"

	"gitlab.com/evatix-go/strhelper/constants"
)

func IntToString(number int) string {
	return fmt.Sprintf(constants.SprintValueFormat, number)
}

func Float32ToString(number float32) string {
	return fmt.Sprintf(constants.SprintValueFormat, number)
}

func Float64ToString(number float64) string {
	return fmt.Sprintf(constants.SprintValueFormat, number)
}

func Int64ToString(number int64) string {
	return fmt.Sprintf(constants.SprintValueFormat, number)
}

func UInt64ToString(number uint64) string {
	return fmt.Sprintf(constants.SprintValueFormat, number)
}

func Int16ToString(number int16) string {
	return fmt.Sprintf(constants.SprintValueFormat, number)
}

func Int8ToString(number int8) string {
	return fmt.Sprintf(constants.SprintValueFormat, number)
}

func UInt8ToString(number uint8) string {
	return fmt.Sprintf(constants.SprintValueFormat, number)
}

// if nil then empty string.
func AnyPtrToString(any *interface{}) string {
	if any == nil || *any == nil {
		return constants.EmptyString
	}

	return fmt.Sprintf(constants.SprintValueFormat, *any)
}

// if nil then empty string.
func AnyToString(any interface{}) string {
	if any == nil {
		return constants.EmptyString
	}

	return fmt.Sprintf(constants.SprintValueFormat, any)
}
