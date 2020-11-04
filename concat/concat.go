package concat

import (
	"fmt"
	"strings"

	"gitlab.com/evatix-go/strhelper"
	"gitlab.com/evatix-go/strhelper/constants"
)

func AnyValues(
	separator string,
	isSkipEmptyOrNil bool,
	contents ...interface{},
) string {
	return anys(
		&separator,
		isSkipEmptyOrNil,
		constants.SprintValueFormat,
		nil,
		&contents)
}

func AnyNameValues(
	separator string,
	isSkipEmptyOrNil bool,
	contents ...interface{},
) string {
	return anys(
		&separator,
		isSkipEmptyOrNil,
		constants.SprintFullPropertyNameValueFormat,
		nil,
		&contents)
}

func AnyFullNameValues(
	separator string,
	isSkipEmptyOrNil bool,
	contents ...interface{},
) string {
	return anys(
		&separator,
		isSkipEmptyOrNil,
		constants.SprintFullPropertyNameValueFormat,
		nil,
		&contents)
}

func Anys(
	separator string,
	isSkipEmptyOrNil bool,
	contentPrintFormat string,
	contents ...interface{},
) string {
	return anys(
		&separator,
		isSkipEmptyOrNil,
		contentPrintFormat,
		nil,
		&contents)
}

func anys(
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

// Empty separator, empty string will be ignored
func Strings(contents ...string) string {
	return strhelper.JoinPtrExceptEmpty(&contents, constants.EmptyStringPtr)
}

// empty string will be ignored
func StringUsingPipe(contents ...string) string {
	return strhelper.JoinPtrExceptEmpty(&contents, constants.PipePtr)
}

// empty string will be ignored
func StringsUsingComma(contents ...string) string {
	return strhelper.JoinPtrExceptEmpty(&contents, constants.CommaPtr)
}

func StringsUsingSpace(contents ...string) string {
	return strhelper.JoinPtrExceptEmpty(&contents, constants.SpacePtr)
}

func StringsUsingHyphen(contents ...string) string {
	return strhelper.JoinPtrExceptEmpty(&contents, constants.HyphenPtr)
}

func StringsWithSeparator(
	currentStr,
	separator string,
	isSkipEmptyOrNil bool,
	contents ...string,
) string {
	return StringsArrayWithSeparator(
		&currentStr,
		&separator,
		isSkipEmptyOrNil,
		&contents)
}

func StringsArrayWithSeparator(
	currentStr,
	separator *string,
	isSkipEmptyOrNil bool,
	contents *[]string,
) string {
	var combinedContents string

	if isSkipEmptyOrNil {
		combinedContents = strhelper.JoinPtrExceptEmpty(contents, separator)
	} else {
		combinedContents = strhelper.JoinPtr(contents, separator)
	}

	if currentStr == nil || *currentStr == constants.EmptyString || len(*currentStr) == 0 {
		return combinedContents
	}

	// TODO this requires optimization
	if combinedContents != constants.EmptyString && len(strings.TrimSpace(combinedContents)) > 0 {
		combinedContents = *separator + combinedContents
	}

	final := *currentStr + combinedContents

	return final
}
