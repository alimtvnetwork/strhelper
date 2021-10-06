package tostr

import (
	"fmt"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/corejson"

	"gitlab.com/evatix-go/strhelper/strhelpercore"
)

func Int(number int) string {
	return fmt.Sprintf(constants.SprintValueFormat, number)
}

func Float32(number float32) string {
	return fmt.Sprintf(constants.SprintValueFormat, number)
}

func Float64(number float64) string {
	return fmt.Sprintf(constants.SprintValueFormat, number)
}

func Int64(number int64) string {
	return fmt.Sprintf(constants.SprintValueFormat, number)
}

func UInt64(number uint64) string {
	return fmt.Sprintf(constants.SprintValueFormat, number)
}

func Int16(number int16) string {
	return fmt.Sprintf(constants.SprintValueFormat, number)
}

func Int8(number int8) string {
	return fmt.Sprintf(constants.SprintValueFormat, number)
}

func UInt8(number uint8) string {
	return fmt.Sprintf(constants.SprintValueFormat, number)
}

// if nil then empty string.
func AnyPtr(any *interface{}) string {
	if any == nil || *any == nil {
		return constants.EmptyString
	}

	return fmt.Sprintf(constants.SprintValueFormat, *any)
}

// if nil then empty string.
// usages constants.SprintValueFormat to print the value of the object.
func Any(any interface{}) string {
	if any == nil {
		return constants.EmptyString
	}

	return fmt.Sprintf(constants.SprintValueFormat, any)
}

func AnyToJsonStrWithErrorPtr(any interface{}) *strhelpercore.StringWithError {
	jsonResult := corejson.NewFromAny(any)

	if jsonResult.HasError() {
		return strhelpercore.NewStringWithErrorOnlyError(jsonResult.MeaningfulError())
	}

	return strhelpercore.NewStringWithNoError(jsonResult.JsonString())
}

// Json
//
// if nil then empty string.
func Json(any interface{}) string {
	jsonResult := AnyToJsonStrWithErrorPtr(&any)

	jsonResult.HandleError()

	return jsonResult.Value()
}
