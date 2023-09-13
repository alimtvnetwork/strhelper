package isstrs

import (
	"gitlab.com/auk-go/strhelper/internal/isanyinternal"
)

// AnyItemsEquals compares leftItems and rightItems and returns bool
//  - If both nil returns true.
//  - If one nil and another is not then returns false.
//  - If both lengths are not same returns false.
//  - If both pointers are same returns true.
//  - If all the items are equals based on encoder encoding to bytes then returns true.
//
// @isContinueOnBothItemParseError
//  - if true then at the same index both item has parse error then continue that means
//      assuming both are same based on error.
//  - if false then at the same index any parse error from binary then returns false no panic.
func AnyItemsEquals(
	leftItems []interface{},
	rightItems []interface{},
	startsAt int,
	isContinueOnBothItemParseError bool,
) bool {
	return isanyinternal.Equals(
		leftItems,
		rightItems,
		startsAt,
		isContinueOnBothItemParseError)
}
