package strconcat

import (
	"gitlab.com/evatix-go/core/constants"

	"gitlab.com/evatix-go/strhelper/whitespace"
)

// Concatenates (@currentStr + *separator + combinedContents) to a single string using @separator.
// @Expression:
//  - @currentStr + *separator + combinedContents
//
// @isSkipEmptyOrNil:
//  - Skip nil or empty string in elements. (not the whitespace)
//  - If final string compiled string from contents is a whitespace then ignored.
//
// @separator:
//  - used to strconcat each strings / elements.
//
// @Returns:
//  - @isSkipEmptyOrNil false , currentStr +  separator + allContents join with separator (skips any with nil or "")
//  - @isSkipEmptyOrNil true ,
//    - if not empty or whitespace (currentStr) then returns allContents join with separator (skips any with nil or "")
//    - if not empty or whitespace (allContents join with separator (skips any with nil or "")) then returns currentStr
//    - if both are not empty and combined contents is not whitespace then returns
//          @currentStr + separator + all contents combined with separator (skips any with nil or "")
func StringsArrayWithSeparator(
	currentStr,
	separator string,
	isSkipEmptyOrNil bool,
	contents *[]string,
) string {
	var combinedContents string

	if isSkipEmptyOrNil {
		combinedContents = JoinPtrExceptEmpty(contents, separator)
	} else {
		combinedContents = JoinPtr(contents, separator)
	}

	if currentStr == constants.EmptyString {
		return combinedContents
	}

	if combinedContents != constants.EmptyString && !whitespace.IsWhitespaces(combinedContents) {
		combinedContents = separator + combinedContents
	}

	return currentStr + combinedContents
}
