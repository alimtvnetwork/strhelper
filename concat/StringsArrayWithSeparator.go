package concat

import (
	"gitlab.com/evatix-go/strhelper/constants"
	"gitlab.com/evatix-go/strhelper/whitespace"
)

// Concatenates the @currentStr with contents to a single string using @separator.
//
// @isSkipEmptyOrNil:
//  - Skip nil or empty string in elements. (not the whitespace)
//  - If final string compiled strings from contents is a whitespace then ignored.
//
// @separator:
//  - used to concat each strings / elements.
//
// Copied from golang library (reference : https://bit.ly/3oPHGdy).
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
