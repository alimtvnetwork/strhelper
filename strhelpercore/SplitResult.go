package strhelpercore

type SplitResult struct {
	SplitPrev string
	// or the splitter
	Separator string
	Index     int
	IsEmpty   bool
}

func (it *SplitResult) IsEquals(another *SplitResult) bool {
	if another == nil && it == nil {
		return true
	}

	if another == nil || it == nil {
		return false
	}

	if another == it {
		return true
	}

	isIndexSame := another.Index == it.Index
	isEmptySame := another.IsEmpty == it.IsEmpty
	if !isEmptySame || !isIndexSame {
		return false
	}

	if it.IsEmpty {
		return true
	}

	isSeparatorSame := it.Separator == another.Separator
	isSplitPrevSame := it.SplitPrev == another.SplitPrev

	return isSeparatorSame && isSplitPrevSame
}
