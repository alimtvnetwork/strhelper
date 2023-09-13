package tostr

import (
	"fmt"
	"reflect"
	"strings"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/converters"
	"gitlab.com/auk-go/core/coredata/corejson"
	"gitlab.com/auk-go/core/errcore"

	"gitlab.com/auk-go/strhelper/strhelpercore"
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
func FromAnyPtr(anyItem *interface{}) string {
	if anyItem == nil || *anyItem == nil {
		return constants.EmptyString
	}

	return fmt.Sprintf(constants.SprintValueFormat, *anyItem)
}

// FromAny
//
// if nil then empty string.
// usages constants.SprintValueFormat to print the value of the object.
func FromAny(anyItem interface{}) string {
	if anyItem == nil {
		return constants.EmptyString
	}

	return fmt.Sprintf(constants.SprintValueFormat, anyItem)
}

func FromAnyToJsonStrWithErrorPtr(anyItem interface{}) *strhelpercore.StringWithError {
	jsonResult := corejson.New(anyItem)

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
func Json(anyItem interface{}) string {
	jsonResult := corejson.New(anyItem)

	if jsonResult.HasError() {
		return jsonResult.MeaningfulErrorMessage()
	}

	return jsonResult.JsonString()
}

// CastingJson
//
// Applies casting
func CastingJson(anyItem interface{}) string {
	jsonResult := corejson.
		AnyTo.
		SerializedJsonResult(anyItem)

	if jsonResult.HasError() {
		return jsonResult.MeaningfulErrorMessage()
	}

	return jsonResult.JsonString()
}

func PrettyJsonStringOrErrString(anyItem interface{}) string {
	jsonResult := corejson.
		AnyTo.
		SerializedJsonResult(anyItem)

	return jsonResult.PrettyJsonStringOrErrString()
}

// SafePrettyJson
//
// if nil then empty string.
func SafePrettyJson(
	anyItem interface{},
) string {
	jsonResult := corejson.New(anyItem)

	return jsonResult.PrettyJsonString()
}

// JsonMust
//
// if nil then empty string.
func JsonMust(anyItem interface{}) string {
	jsonResult := corejson.New(anyItem)
	jsonResult.HandleError()

	return jsonResult.JsonString()
}

// AnyItemOption
//
//  Full Fields : %#v
//  Value       : %v
func AnyItemOption(
	isFields bool,
	anyItem interface{},
) string {
	return converters.Any.ToString(
		isFields,
		anyItem)
}

// AnyItemFullString
//
//  %#v
func AnyItemFullString(
	anyItem interface{},
) string {
	return converters.Any.ToString(
		true,
		anyItem)
}

// AnyItemHashFullString
//
//  %#v
func AnyItemHashFullString(
	anyItem interface{},
) string {
	return converters.Any.ToString(
		true,
		anyItem)
}

// AnyItemFullStringValue
//
//  %+v
func AnyItemFullStringValue(
	anyItem interface{},
) string {
	return converters.Any.FullString(
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

func BytesPtr(
	rawBytes *[]byte,
) string {
	if rawBytes == nil || len(*rawBytes) == 0 {
		return ""
	}

	return string(*rawBytes)
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
