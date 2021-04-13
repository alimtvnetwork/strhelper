package strconcat

import (
	"gitlab.com/evatix-go/core/constants"
)

// Concatenates the (@preContents to a string using separator) with
// (@postContents to a string using separator) using separator to a string.
//
// Expression:
//  - (@preContents to a string using separator) + separator + (@postContents to a string using separator)
//
// @isSkipEmptyOrNil:
//  - Skip nil or empty string in elements. (not the whitespace)
//  - If final string compiled string from preContents or postContents then it is ignored.
//
// @separator:
//  - used to strconcat each strings / elements.
//
// @Returns:
//  - @isSkipEmptyOrNil false , (preContents joined with separator) + separator + (postContents joined with separator)
//  - @isSkipEmptyOrNil true ,
//    - if empty (preContents combined) then returns combined [postContents] combined with separator.
//    - if empty (postContents combined) then returns combined [preContents] combined with separator.
//    - if both are not empty and combined string is not whitespace then
//          (preContents joined with separator) + separator + (postContents joined with separator)
func CombineArrayWithAnotherUsingSeparator(
	separator string,
	isSkipEmptyOrNil bool,
	preContents *[]string,
	postContents *[]string,
) string {
	var preContentsCombined, postContentsCombined string

	if isSkipEmptyOrNil {
		preContentsCombined = JoinPtrExceptEmpty(preContents, separator)
	} else {
		preContentsCombined = JoinPtr(preContents, separator)
	}

	if isSkipEmptyOrNil {
		postContentsCombined = JoinPtrExceptEmpty(postContents, separator)
	} else {
		postContentsCombined = JoinPtr(postContents, separator)
	}

	if preContentsCombined == constants.EmptyString || len(preContentsCombined) == 0 {
		return postContentsCombined
	}

	if postContentsCombined == constants.EmptyString || len(postContentsCombined) == 0 {
		return preContentsCombined
	}

	return preContentsCombined + separator + postContentsCombined
}
