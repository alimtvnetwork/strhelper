package isanyinternal

import "gitlab.com/auk-go/strhelper/internal/coreinternal"

// ItemsNullDeduction compares leftItems and rightItems and
// returns coreinternal.BoolResultWrapper
//  - If both nil returns true.
//  - If one nil and another is not then returns false.
//  - If both pointers are same returns true.
//  - If none of the conditions satisfied then returns coreinternal.NewBoolResultWrapperNotApplicable()
func ItemsNullDeduction(
	leftItems *[]interface{},
	rightItems *[]interface{},
) coreinternal.BoolResultWrapper {
	// if pointer same
	if leftItems == rightItems {
		return coreinternal.NewBoolResultWrapperTrue()
	}

	if leftItems == nil && rightItems == nil {
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

	return coreinternal.NewBoolResultWrapperNotApplicable()
}
