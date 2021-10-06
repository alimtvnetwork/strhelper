package strconcat

import (
	"gitlab.com/evatix-go/strhelper/ds/strhashset"
	"gitlab.com/evatix-go/strhelper/internal/panichelper"
	"gitlab.com/evatix-go/strhelper/strs/isstrs"
)

// ArraysOfArraysToArray
//
// Concatenates the arrays of array items to a single array.
//
// @Expression:
//  - Generate capacity first by loop through all the @linesArray
//  - Create a slice based on the capacity found.
//  - start copying item from array of array if meets the condition
//  ((isSkipEmptyOrNil && not nil) && Not found in skipFilter)
//
// @isSkipEmptyOrNil:
//  - Skip nil or empty string in elements. (not the whitespace)
//
// @Returns:
//  - @isSkipEmptyOrNil false , all will be added to slice. If single one is nil panic.
//  - @isSkipEmptyOrNil true ,
//    - if nil array will be skipped
//    - if line in array if empty or nil will be skipped
func ArraysOfArraysToArray(
	isSkipEmptyOrNil bool,
	skipFilter *strhashset.Hashset,
	linesArray ...*[]string,
) *[]string {
	if skipFilter == nil {
		skipFilter = strhashset.NewEmpty()
	}

	if linesArray == nil {
		panichelper.NullReference("linesArray")
	}

	capacity := 0
	for _, array := range linesArray {
		if isSkipEmptyOrNil && array == nil {
			continue
		} else if !isSkipEmptyOrNil && array == nil {
			panichelper.NullReferenceMsg("Cannot have any nil in linesArray", "linesArray")
		}

		//goland:noinspection GoNilness
		capacity += len(*array)
	}

	if skipFilter.IsEmptySet() {
		return copyingArraysOfArrayToArrayNoFilter(
			isSkipEmptyOrNil,
			&linesArray,
			capacity)
	}

	return copyingArraysOfArrayToArrayWithFilter(
		isSkipEmptyOrNil,
		&linesArray,
		capacity,
		skipFilter)
}

func copyingArraysOfArrayToArrayNoFilter(
	isSkipEmptyOrNil bool,
	linesArray *[]*[]string,
	capacity int,
) *[]string {
	// nothing to filter for
	finalLines := make([]string, 0, capacity)
	for _, lines := range *linesArray {
		if isSkipEmptyOrNil && isstrs.Empty(lines) {
			continue
		}

		for _, line := range *lines {
			if isSkipEmptyOrNil && line == "" {
				continue
			}

			finalLines = append(finalLines, line)
		}
	}

	return &finalLines
}

func copyingArraysOfArrayToArrayWithFilter(
	isSkipEmptyOrNil bool,
	linesArray *[]*[]string,
	capacity int,
	skipFilter *strhashset.Hashset,
) *[]string {
	finalLines := make([]string, 0, capacity)
	for _, lines := range *linesArray {
		if isSkipEmptyOrNil && isstrs.Empty(lines) {
			continue
		}

		for _, line := range *lines {
			if (isSkipEmptyOrNil && line == "") || skipFilter.Has(line) {
				continue
			}

			finalLines = append(finalLines, line)
		}
	}

	return &finalLines
}
