package isstrs

import (
	"gitlab.com/evatix-go/strhelper/internal/pkg/isanyinternal"
)

// AnyPointersEquals compares leftItems and rightItems and returns bool
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
func AnyPointersEquals(
	leftItems *[]*interface{},
	rightItems *[]*interface{},
	startsAt int,
	isContinueOnBothItemParseError bool,
) bool {
	return isanyinternal.PointersOfPointersAnyItemsEquals(
		leftItems,
		rightItems,
		startsAt,
		isContinueOnBothItemParseError)
}
