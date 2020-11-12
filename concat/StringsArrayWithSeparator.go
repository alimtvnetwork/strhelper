package concat

import (
	"gitlab.com/evatix-go/strhelper/constants"
	"gitlab.com/evatix-go/strhelper/whitespace"
)

func StringsArrayWithSeparator(
	currentStr,
	separator *string,
	isSkipEmptyOrNil bool,
	contents *[]string,
) string {
	var combinedContents string

	if isSkipEmptyOrNil {
		combinedContents = JoinPtrExceptEmpty(contents, separator)
	} else {
		combinedContents = JoinPtr(contents, separator)
	}

	if currentStr == nil || *currentStr == constants.EmptyString || len(*currentStr) == 0 {
		return combinedContents
	}

	if combinedContents != constants.EmptyString && !whitespace.IsWhitespaces(&combinedContents) {
		combinedContents = *separator + combinedContents
	}

	return *currentStr + combinedContents
}
