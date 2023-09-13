package isanyinternal

import (
	"bytes"

	"gitlab.com/auk-go/strhelper/anyto"
)

// ItemsPointerEquals compares leftItems and rightItems and returns bool
//  - If both nil returns true.
//  - If one nil and another is not then returns false.
//  - If both lengths are not same returns false.
//  - If both pointers are same returns true.
//  - If all the items are equals based on encoder encoding to bytes then returns true.
//
// @isContinueOnBothItemParseError
//  - if true then if at the same index both item has parse error then continue that means
//      assuming both are same based on error.
//  - if false then if at the same index any parse error from binary then returns false no panic.
func ItemsPointerEquals(
	leftItems *[]interface{},
	rightItems *[]interface{},
	startsAt int,
	isContinueOnBothItemParseError bool,
) bool {
	resultWrapper := ItemsNullDeduction(leftItems, rightItems)
	if resultWrapper.IsApplicable {
		return resultWrapper.Result
	}

	leftLength := len(*leftItems)
	rightLength := len(*rightItems)

	if leftLength != rightLength {
		return false
	}

	if startsAt < 0 || startsAt > (leftLength-1) {
		panic("Start cannot be negative or larger than length index.")
	}

	for ; startsAt < leftLength; startsAt++ {
		left := (*leftItems)[startsAt]
		right := (*rightItems)[startsAt]

		if left == nil && right == nil {
			continue
		}

		if left == nil || right == nil {
			return false
		}

		leftBytes, lError := anyto.Bytes(left)
		rightBytes, rError := anyto.Bytes(right)

		if isContinueOnBothItemParseError && lError != nil && rError != nil {
			continue
		}

		if !isContinueOnBothItemParseError && (lError != nil || rError != nil) {
			return false
		}

		if !bytes.Equal(leftBytes, rightBytes) {
			return false
		}
	}

	return true
}
