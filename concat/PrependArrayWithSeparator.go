package concat

import (
	"gitlab.com/evatix-go/strhelper/strconst"
	"gitlab.com/evatix-go/strhelper/whitespace"
)

// Concatenates (combinedContents + *separator + @currentStr) to a single string using @separator.
//
// @Expression:
//  - combinedContents + *separator + @currentStr
// @isSkipEmptyOrNil:
//  - Skip nil or empty string in elements. (not the whitespace)
//  - If final string compiled string from contents is a whitespace then ignored.
//
// @separator:
//  - used to concat each strings / elements.
//
// @Returns:
//  - @isSkipEmptyOrNil false , (allContents joined with separator) + separator + currentStr
//  - @isSkipEmptyOrNil true ,
//    - if not empty or whitespace (currentStr) then returns allContents join with separator (skips any with nil or "")
//    - if not empty or whitespace (allContents join with separator(skips any with nil or "")) then returns currentStr
//    - if both are not empty and combined contents is not whitespace then returns (all contents combined with separator (skips any with nil or "")) + separator + @currentStr
func PrependArrayWithSeparator(
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

	if currentStr == nil || *currentStr == strconst.EmptyString || len(*currentStr) == 0 {
		return combinedContents
	}

	if combinedContents != strconst.EmptyString && !whitespace.IsWhitespaces(&combinedContents) {
		combinedContents = combinedContents + *separator
	}

	return combinedContents + *currentStr
}
