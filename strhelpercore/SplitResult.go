package strhelpercore

import "gitlab.com/evatix-go/strhelper/internal/isstrinternal"

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

	isSplitPrevNull := isstrinternal.PointerEqualBasedOnAddressDeduction(splitResult.SplitPrev, another.SplitPrev)

	if isSplitPrevNull.IsApplicableWithFalse() {
		return false
	}

	isSeparatorNull := isstrinternal.PointerEqualBasedOnAddressDeduction(splitResult.Separator, another.Separator)

	if isSeparatorNull.IsApplicableWithFalse() {
		return false
	}

	isSeparatorSame := isstrinternal.EqualsPtr(splitResult.Separator, another.Separator)
	isSplitPrevSame := isstrinternal.EqualsPtr(splitResult.SplitPrev, another.SplitPrev)

	return isSeparatorSame && isSplitPrevSame
}
