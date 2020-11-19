package concat

import (
	"gitlab.com/evatix-go/strhelper/strconst"
)

// Concatenates the (@preContents to a string using separator) with (@postContents to a string using separator) using separator to a string.
//
// Expression:
//  - join all many of manyArrays to a single string using separator.
//
// @isSkipEmptyOrNil:
//  - Skip on nil or empty string in elements. (not the whitespace)
//  - If final string compiled string from contents is a whitespace then ignored.
//
// @separator:
//  - used to concat each strings / elements.
//
// @Returns:
//  - @isSkipEmptyOrNil false , use separator in between contents for even empty or nil string
//  - @isSkipEmptyOrNil true ,
//    - concatenates only if the content is not nil or empty. (if has any empty space " " will consider not empty)
func ManyArraysCollectionUsingSeparator(
	separator *string,
	isSkipEmptyOrNil bool,
	manyArrays *[]*[]string,
) string {
	length := len(*manyArrays)

	if *manyArrays == nil || length == 0 {
		return strconst.EmptyString
	}

	finalProcessedArray := make([]*string, 0, length+1)
	if isSkipEmptyOrNil {
		for _, singleStringArray := range *manyArrays {
			currentCombined := JoinPtrExceptEmpty(singleStringArray, separator)

			if currentCombined != "" {
				finalProcessedArray = append(finalProcessedArray, &currentCombined)
			}
		}
	} else {
		for _, singleStringArray := range *manyArrays {
			currentCombined := JoinPtr(singleStringArray, separator)
			finalProcessedArray = append(finalProcessedArray, &currentCombined)
		}
	}

	if isSkipEmptyOrNil {
		return JoinStrPtrExceptEmpty(&finalProcessedArray, separator)
	} else {
		return JoinStrPtr(&finalProcessedArray, separator)
	}
}
