package strconcat

import (
	"fmt"

	"gitlab.com/evatix-go/core/constants"
)

// Concat any objects to single string using sprintf format given constants.SprintValueFormat
//
// @isSkipEmptyOrNil
//  - enabled : Skips singleContent or any contents in content if nil.
//
// @separator:
//  - it is used to strconcat each contents to string to single one.
func AnyValues(
	separator string,
	isSkipEmptyOrNil bool,
	contents ...interface{},
) string {
	return AnyArrayOfInterfaces(
		separator,
		isSkipEmptyOrNil,
		constants.SprintValueFormat,
		nil,
		&contents)
}

// Concat any objects to single string using sprintf format given constants.SprintPropertyNameValueFormat
//
// @isSkipEmptyOrNil
//  - enabled : Skips singleContent or any contents in content if nil.
//
// @separator:
//  - it is used to strconcat each contents to string to single one.
func AnyNameValues(
	separator string,
	isSkipEmptyOrNil bool,
	contents ...interface{},
) string {
	return AnyArrayOfInterfaces(
		separator,
		isSkipEmptyOrNil,
		constants.SprintPropertyNameValueFormat,
		nil,
		&contents)
}

// Concat any objects to single string using sprintf format given constants.SprintFullPropertyNameValueFormat
//
// @isSkipEmptyOrNil
//  - enabled : Skips singleContent or any contents in content if nil.
//
// @separator:
//  - it is used to strconcat each contents to string to single one.
func AnyFullNameValues(
	separator string,
	isSkipEmptyOrNil bool,
	contents ...interface{},
) string {
	return AnyArrayOfInterfaces(
		separator,
		isSkipEmptyOrNil,
		constants.SprintFullPropertyNameValueFormat,
		nil,
		&contents)
}

// AnyItems
//
// Concat any objects to single string using its sprintf format given
//
// @isSkipEmptyOrNil
//  - enabled : Skips singleContent or any contents in content if nil.
//
// @separator:
//  - it is used to string concat each contents to string to single one.
func AnyItems(
	separator string,
	isSkipEmptyOrNil bool,
	contentPrintFormat string,
	contents ...interface{},
) string {
	return AnyArrayOfInterfaces(
		separator,
		isSkipEmptyOrNil,
		contentPrintFormat,
		nil,
		&contents)
}

// AnyArrayOfInterfaces
//
// Concat any objects to single string using its sprintf format given
//
// @isSkipEmptyOrNil
//  - enabled : Skips singleContent or any contents in content if nil.
//
// @separator:
//  - it is used to string concat each contents to string to single one.
func AnyArrayOfInterfaces(
	separator string,
	isSkipEmptyOrNil bool,
	contentPrintFormat string,
	singleContent *interface{},
	contents *[]interface{},
) string {
	newLines := make([]string, 0, len(*contents)+2)
	firstLine := ""

	if singleContent != nil {
		firstLine = fmt.Sprintf(contentPrintFormat, singleContent)
	}

	if isSkipEmptyOrNil && len(firstLine) > 0 {
		newLines = append(newLines, firstLine)
	} else if isSkipEmptyOrNil == false {
		newLines = append(newLines, constants.NilString)
	}

	for _, content := range *contents {
		if isSkipEmptyOrNil && content == nil {
			continue
		}

		newLines = append(
			newLines,
			fmt.Sprintf(contentPrintFormat, content))
	}

	return StringsArrayWithSeparator(
		constants.EmptyString,
		separator,
		isSkipEmptyOrNil,
		&newLines)
}

// Concat any objects to string using compiler function.
//
// @isSkipEmptyOrNil
//  - enabled : Skips singleContent or any contents in content if nil.
//
// @separator:
//  - it is used to strconcat each lines. Whereas compiler function only compile the single interface to string only.
func AnyArrayOfInterfacesUsingFunc(
	separator string,
	isSkipEmptyOrNil bool,
	singleContent *interface{},
	contents *[]interface{},
	compiler func(any interface{}) string,
) string {
	newLines := make([]string, 0, len(*contents)+2)
	firstLine := ""

	if singleContent != nil {
		firstLine = compiler(singleContent)
	}

	if isSkipEmptyOrNil && len(firstLine) > 0 {
		newLines = append(newLines, firstLine)
	} else if isSkipEmptyOrNil == false {
		newLines = append(newLines, constants.NilString)
	}

	for _, content := range *contents {
		if isSkipEmptyOrNil && content == nil {
			continue
		}

		newLines = append(newLines, compiler(content))
	}

	return StringsArrayWithSeparator(
		constants.EmptyString,
		separator,
		isSkipEmptyOrNil,
		&newLines)
}
