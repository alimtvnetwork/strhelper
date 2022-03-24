package tostr

import (
	"fmt"
	"reflect"
	"strings"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/converters"
	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/core/errcore"

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

// AnyPtr
//
//  if nil then empty string.
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
	jsonResult := corejson.New(any)

	if jsonResult.HasError() {
		return strhelpercore.NewStringWithErrorOnlyError(jsonResult.MeaningfulError())
	}

	return strhelpercore.NewStringWithNoError(jsonResult.JsonString())
}

// Json
//
// if nil then empty string.
//
// On error returns error as string.
func Json(any interface{}) string {
	jsonResult := AnyToJsonStrWithErrorPtr(&any)

	if jsonResult.HasError() {
		return jsonResult.Error().Error()
	}

	return jsonResult.Value()
}

// JsonMust
//
// if nil then empty string.
func JsonMust(any interface{}) string {
	jsonResult := AnyToJsonStrWithErrorPtr(&any)
	jsonResult.HandleError()

	return jsonResult.Value()
}

func AnyItemOption(
	isFields bool,
	anyItem interface{},
) string {
	return converters.AnyToString(
		isFields,
		anyItem)
}

func Bytes(
	rawBytes []byte,
) string {
	if len(rawBytes) == 0 {
		return ""
	}

	return string(rawBytes)
}

func AnyItemWithFields(
	anyItem interface{},
) string {
	return converters.AnyToString(
		true,
		anyItem)
}

func TypeNameOption(
	isSafeChecking bool,
	anyItem interface{},
) string {
	if isSafeChecking {
		return SafeTypeName(anyItem)
	}

	return reflect.TypeOf(anyItem).String()
}

func SafeTypeName(
	anyItem interface{},
) string {
	rf := reflect.TypeOf(anyItem)

	if rf == nil {
		return ""
	}

	return rf.String()
}

func Error(
	err error,
) string {
	return errcore.ToString(err)
}

// FromLines
//
//  join using constants.DefaultLine
func FromLines(
	lines ...string,
) string {
	return strings.Join(
		lines,
		constants.DefaultLine)
}
