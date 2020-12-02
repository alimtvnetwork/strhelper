package isinternal

import "gitlab.com/evatix-go/strhelper/internal/coreinternal"

// PointerEqualBasedOnAddressDeduction compares leftItems and rightItems and returns coreinternal.BoolResultWrapper
//  - If both nil returns true.
//  - If one nil and another is not then returns false.
//  - If both pointers are same returns true.
//  - If none of the conditions satisfied then returns coreinternal.NewBoolResultWrapperNotApplicable()
func PointerEqualBasedOnAddressDeduction(
	left *string,
	right *string,
) coreinternal.BoolResultWrapper {
	if left == right && left == nil {
		return coreinternal.NewBoolResultWrapperTrue()
	}

	if left == nil || right == nil {
		return coreinternal.NewBoolResultWrapperFalse()
	}

	// if pointer same
	if left == right {
		return coreinternal.NewBoolResultWrapperTrue()
	}

	return coreinternal.NewBoolResultWrapperNotApplicable()
}
