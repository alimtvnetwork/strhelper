package concat

import (
	"fmt"

	"gitlab.com/evatix-go/strhelper/constants"
)

func AnyValues(
	separator string,
	isSkipEmptyOrNil bool,
	contents ...interface{},
) string {
	return AnyArrayOfInterfaces(
		&separator,
		isSkipEmptyOrNil,
		constants.SprintValueFormat,
		nil,
		&contents)
}

// Concat any object to string, sprintf format given constants.SprintPropertyNameValueFormat
func AnyNameValues(
	separator string,
	isSkipEmptyOrNil bool,
	contents ...interface{},
) string {
	return AnyArrayOfInterfaces(
		&separator,
		isSkipEmptyOrNil,
		constants.SprintPropertyNameValueFormat,
		nil,
		&contents)
}

// Concat any object to string, sprintf format given constants.SprintFullPropertyNameValueFormat
func AnyFullNameValues(
	separator string,
	isSkipEmptyOrNil bool,
	contents ...interface{},
) string {
	return AnyArrayOfInterfaces(
		&separator,
		isSkipEmptyOrNil,
		constants.SprintFullPropertyNameValueFormat,
		nil,
		&contents)
}

// Concat any object to string using it's sprintf format given
func Anys(
	separator string,
	isSkipEmptyOrNil bool,
	contentPrintFormat string,
	contents ...interface{},
) string {
	return AnyArrayOfInterfaces(
		&separator,
		isSkipEmptyOrNil,
		contentPrintFormat,
		nil,
		&contents)
}

// Concat any object to string using it's sprintf format given
func AnyArrayOfInterfaces(
	separator *string,
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
	} else {
		newLines = append(newLines, constants.NilString)
	}

	for _, content := range *contents {
		if isSkipEmptyOrNil && content == nil {
			continue
		}

		newLines = append(newLines, fmt.Sprintf(contentPrintFormat, content))
	}

	return StringsArrayWithSeparator(nil, separator, isSkipEmptyOrNil, &newLines)
}

// Concat any object to string using it's sprintf format given
func AnyArrayOfInterfacesUsingFunc(
	separator *string,
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
	} else {
		newLines = append(newLines, constants.NilString)
	}

	for _, content := range *contents {
		if isSkipEmptyOrNil && content == nil {
			continue
		}

		newLines = append(newLines, compiler(content))
	}

	return StringsArrayWithSeparator(nil, separator, isSkipEmptyOrNil, &newLines)
}
