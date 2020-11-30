package tostr

import (
	"encoding/json"
	"fmt"

	"gitlab.com/evatix-go/strhelper/strconst"
	"gitlab.com/evatix-go/strhelper/strhelpercore"
)

func Int(number int) string {
	return fmt.Sprintf(strconst.SprintValueFormat, number)
}

func Float32(number float32) string {
	return fmt.Sprintf(strconst.SprintValueFormat, number)
}

func Float64(number float64) string {
	return fmt.Sprintf(strconst.SprintValueFormat, number)
}

func Int64(number int64) string {
	return fmt.Sprintf(strconst.SprintValueFormat, number)
}

func UInt64(number uint64) string {
	return fmt.Sprintf(strconst.SprintValueFormat, number)
}

func Int16(number int16) string {
	return fmt.Sprintf(strconst.SprintValueFormat, number)
}

func Int8(number int8) string {
	return fmt.Sprintf(strconst.SprintValueFormat, number)
}

func UInt8(number uint8) string {
	return fmt.Sprintf(strconst.SprintValueFormat, number)
}

// if nil then empty string.
func AnyPtr(any *interface{}) string {
	if any == nil || *any == nil {
		return strconst.EmptyString
	}

	return fmt.Sprintf(strconst.SprintValueFormat, *any)
}

// if nil then empty string.
// usages strconst.SprintValueFormat to print the value of the object.
func Any(any interface{}) string {
	if any == nil {
		return strconst.EmptyString
	}

	return fmt.Sprintf(strconst.SprintValueFormat, any)
}

func AnyToJsonStrWithErrorPtr(any *interface{}) *strhelpercore.StringWithError {
	jsonBytes, er := json.Marshal(any)

	if er != nil {
		return strhelpercore.NewStringWithErrorOnlyError(&er)
	}

	jsonStr := string(jsonBytes)

	return strhelpercore.NewStringWithNoError(&jsonStr)
}

// if nil then empty string.
func Json(any interface{}) string {
	jsonResult := AnyToJsonStrWithErrorPtr(&any)

	jsonResult.HasError()

	return *jsonResult.Value()
}
