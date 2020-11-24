package strhelpercore

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

// NewBoolResultWrapper returns new
func NewBoolResultWrapper(result bool) BoolResultWrapper {
	return BoolResultWrapper{Result: result, IsApplicable: true}
}

func NewBoolResultWrapperNotApplicable() BoolResultWrapper {
	return BoolResultWrapper{Result: false, IsApplicable: false}
}
