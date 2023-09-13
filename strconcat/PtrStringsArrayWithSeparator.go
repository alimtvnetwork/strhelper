package strconcat

import (
	"gitlab.com/auk-go/core/constants"

	"gitlab.com/auk-go/strhelper/whitespace"
)

// Concatenates the contents and @currentStr to a single string using @separator.
//
// @isSkipEmptyOrNil:
//  - Skip nil or empty string in elements. (not the whitespace)
//  - If final string compiled strings from contents is a whitespace then ignored.
//
// @separator:
//  - used to strconcat each strings / elements.
//
// Copied from golang library (reference : https://bit.ly/3oPHGdy).
func PtrStringsArrayWithSeparator(
	currentStr,
	separator string,
	isSkipEmptyOrNil bool,
	contents []*string,
) string {
	var combinedContents string

	if isSkipEmptyOrNil {
		combinedContents = JoinStrPtrExceptEmpty(contents, separator)
	} else {
		combinedContents = JoinStrPtr(contents, separator)
	}

	if currentStr == constants.EmptyString {
		return combinedContents
	}

	if combinedContents != constants.EmptyString && !whitespace.IsWhitespaces(combinedContents) {
		combinedContents = separator + combinedContents
	}

	return currentStr + combinedContents
}
