package strhelpercore

import "gitlab.com/evatix-go/strhelper/internal/isinternal"

type SplitResult struct {
	SplitPrev *string
	// or the splitter
	Separator *string
	Index     int
	IsEmpty   bool
}

func (splitResult *SplitResult) IsEquals(another *SplitResult) bool {
	if another == nil {
		return false
	}

	isIndexSame := another.Index == splitResult.Index
	isEmptySame := another.IsEmpty == splitResult.IsEmpty
	if !isEmptySame || !isIndexSame {
		return false
	}

	if splitResult.IsEmpty {
		return true
	}

	isSplitPrevNull := isinternal.PointerEqualBasedOnAddressDeduction(splitResult.SplitPrev, another.SplitPrev)

	if isSplitPrevNull.IsApplicableWithFalse() {
		return false
	}

	isSeparatorNull := isinternal.PointerEqualBasedOnAddressDeduction(splitResult.Separator, another.Separator)

	if isSeparatorNull.IsApplicableWithFalse() {
		return false
	}

	isSeparatorSame := isinternal.EqualsPtr(splitResult.Separator, another.Separator)
	isSplitPrevSame := isinternal.EqualsPtr(splitResult.SplitPrev, another.SplitPrev)

	return isSeparatorSame && isSplitPrevSame
}
