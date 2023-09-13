package isanyinternal

import "gitlab.com/auk-go/strhelper/internal/coreinternal"

// NullDeduction compares leftItems and rightItems and returns coreinternal.BoolResultWrapper
//  - If both nil returns true.
//  - If one nil and another is not then returns false.
//  - If both pointers are same returns true.
//  - If none of the conditions satisfied then returns coreinternal.NewBoolResultWrapperNotApplicable()
func NullDeduction(
	left interface{},
	right interface{},
) coreinternal.BoolResultWrapper {
	// is pointer same
	if left == right {
		return coreinternal.NewBoolResultWrapperTrue()
	}

	if left == nil && right == nil {
		return coreinternal.NewBoolResultWrapperTrue()
	}

	if left == nil || right == nil {
		return coreinternal.NewBoolResultWrapperFalse()
	}

	return coreinternal.NewBoolResultWrapperNotApplicable()
}
