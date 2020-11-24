package isstrsinternal

import "gitlab.com/evatix-go/strhelper/internal/pkg/coreinternal"

// isAnyEqualBasedOnAddressDeduction compares leftItems and rightItems and returns coreinternal.BoolResultWrapper
//  - If both nil returns true.
//  - If one nil and another is not then returns false.
//  - If both pointers are same returns true.
//  - If none of the conditions satisfied then returns coreinternal.NewBoolResultWrapperNotApplicable()
func isAnyEqualBasedOnAddressDeduction(
	leftItems *[]interface{},
	rightItems *[]interface{},
) coreinternal.BoolResultWrapper {
	if leftItems == rightItems && leftItems == nil {
		return coreinternal.NewBoolResultWrapperTrue()
	}

	if leftItems == nil || rightItems == nil {
		return coreinternal.NewBoolResultWrapperFalse()
	}

	if *leftItems == nil && *rightItems == nil {
		return coreinternal.NewBoolResultWrapperTrue()
	}

	if *leftItems == nil || *rightItems == nil {
		return coreinternal.NewBoolResultWrapperFalse()
	}

	// if pointer same
	if leftItems == rightItems {
		return coreinternal.NewBoolResultWrapperTrue()
	}

	return coreinternal.NewBoolResultWrapperNotApplicable()
}
