package strconcat

import (
	"gitlab.com/auk-go/core/constants"

	"gitlab.com/auk-go/strhelper/whitespace"
)

// PrependArrayWithCurrentStringUsingSeparator : Prepends array contents before
// @currentStr (@combinedContents + *separator + @currentStr) and
// compiles to a single string using @separator.
//
// @Expression:
//  - @combinedContents = contents joined to single one using separator (skip empty if flag is enabled)
//  - returns @combinedContents + *separator + @currentStr
//
// @isSkipEmptyOrNil:
//  - Skip nil or empty string in elements. (not the whitespace)
//  - If final string compiled string from contents is a whitespace then ignored.
//
// @separator:
//  - used to strconcat each strings / elements.
//
// @Returns:
//  - @isSkipEmptyOrNil false , (allContents joined with separator) + separator + currentStr
//  - @isSkipEmptyOrNil true ,
//    - if not empty or whitespace (currentStr) then returns @combinedContents
//    - if not empty or whitespace (@combinedContents) then returns currentStr
//    - if both are not empty and combined contents is not whitespace then returns
//          (@combinedContents) + separator + @currentStr
func PrependArrayWithCurrentStringUsingSeparator(
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
		combinedContents = combinedContents + separator
	}

	return combinedContents + currentStr
}
