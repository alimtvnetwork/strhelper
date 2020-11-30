package coreinternal

// Some functions usages two return values, one is the result, another if the function condition satisfied.
// Data size example : https://play.golang.org/p/gbVsdr3DB6I
type BoolResultWrapper struct {
	// Function result.
	Result bool
	// If any of the function conditions satisfied.
	//
	// If this is false then Result is invalid.
	IsApplicable bool
}

// NewBoolResultWrapper returns new {Result: false, IsApplicable: true}
func NewBoolResultWrapperFalse() BoolResultWrapper {
	return BoolResultWrapper{Result: false, IsApplicable: true}
}

// NewBoolResultWrapper returns new {Result: true, IsApplicable: true}
func NewBoolResultWrapperTrue() BoolResultWrapper {
	return BoolResultWrapper{Result: true, IsApplicable: true}
}

// NewBoolResultWrapper returns new {Result: result, IsApplicable: true}
func NewBoolResultWrapper(result bool) BoolResultWrapper {
	return BoolResultWrapper{Result: result, IsApplicable: true}
}

func NewBoolResultWrapperNotApplicable() BoolResultWrapper {
	return BoolResultWrapper{Result: false, IsApplicable: false}
}

func (boolResultWrapper *BoolResultWrapper) IsApplicableWithTrue() bool {
	return boolResultWrapper.IsApplicable && boolResultWrapper.Result
}

func (boolResultWrapper *BoolResultWrapper) IsApplicableWithFalse() bool {
	return boolResultWrapper.IsApplicable && !boolResultWrapper.Result
}
