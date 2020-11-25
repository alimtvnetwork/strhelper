package isanyinternal

import "gitlab.com/evatix-go/strhelper/internal/pkg/coreinternal"

// PointerEqualBasedOnAddressDeduction compares leftItems and rightItems and returns coreinternal.BoolResultWrapper
//  - If both nil returns true.
//  - If one nil and another is not then returns false.
//  - If both pointers are same returns true.
//  - If none of the conditions satisfied then returns coreinternal.NewBoolResultWrapperNotApplicable()
func PointerEqualBasedOnAddressDeduction(
	left *interface{},
	right *interface{},
) coreinternal.BoolResultWrapper {
	// if pointer same
	if left == right {
		return coreinternal.NewBoolResultWrapperTrue()
	}

	if left == nil && right == nil {
		return coreinternal.NewBoolResultWrapperTrue()
	}

	if left == nil || right == nil {
		return coreinternal.NewBoolResultWrapperFalse()
	}

	if *left == nil && *right == nil {
		return coreinternal.NewBoolResultWrapperTrue()
	}

	if *left == nil || *right == nil {
		return coreinternal.NewBoolResultWrapperFalse()
	}

	return coreinternal.NewBoolResultWrapperNotApplicable()
}
