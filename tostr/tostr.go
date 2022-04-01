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

func FromInt(number int) string {
	return fmt.Sprintf(constants.SprintValueFormat, number)
}

func FromFloat32(number float32) string {
	return fmt.Sprintf(constants.SprintValueFormat, number)
}

func FromFloat64(number float64) string {
	return fmt.Sprintf(constants.SprintValueFormat, number)
}

func FromInt64(number int64) string {
	return fmt.Sprintf(constants.SprintValueFormat, number)
}

func FromUInt64(number uint64) string {
	return fmt.Sprintf(constants.SprintValueFormat, number)
}

func FromInt16(number int16) string {
	return fmt.Sprintf(constants.SprintValueFormat, number)
}

func FromInt8(number int8) string {
	return fmt.Sprintf(constants.SprintValueFormat, number)
}

func FromUInt8(number uint8) string {
	return fmt.Sprintf(constants.SprintValueFormat, number)
}

// FromAnyPtr
//
//  if nil then empty string.
func FromAnyPtr(any *interface{}) string {
	if any == nil || *any == nil {
		return constants.EmptyString
	}

	return fmt.Sprintf(constants.SprintValueFormat, *any)
}

// if nil then empty string.
// usages constants.SprintValueFormat to print the value of the object.
func FromAny(any interface{}) string {
	if any == nil {
		return constants.EmptyString
	}

	return fmt.Sprintf(constants.SprintValueFormat, any)
}

func FromAnyToJsonStrWithErrorPtr(any interface{}) *strhelpercore.StringWithError {
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
	jsonResult := corejson.New(any)

	if jsonResult.HasError() {
		return jsonResult.MeaningfulErrorMessage()
	}

	return jsonResult.JsonString()
}

// SafePrettyJson
//
// if nil then empty string.
func SafePrettyJson(any interface{}) string {
	jsonResult := corejson.New(any)

	return jsonResult.PrettyJsonString()
}

// JsonMust
//
// if nil then empty string.
func JsonMust(any interface{}) string {
	jsonResult := corejson.New(any)
	jsonResult.HandleError()

	return jsonResult.JsonString()
}

func AnyItemOption(
	isFields bool,
	anyItem interface{},
) string {
	return converters.Any.ToString(
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
	return converters.Any.ToString(
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

func FromPointer(
	pointerString *string,
) string {
	if pointerString == nil {
		return constants.EmptyString
	}

	return *pointerString
}

func FromPointerUsingDefault(
	defaultVal string,
	pointerString *string,
) (output string, isDefault bool) {
	if pointerString == nil {
		return defaultVal, true
	}

	return *pointerString, false
}
