package strhelper

// MergeSlicesPtrOfSlices Don't include nil or length 0 slices
func MergeSlicesPtrOfSlices(slices ...*[]string) *[]string {
	sliceLength := len(slices)

	if sliceLength == 0 {
		return &[]string{}
	}

	return MergeSlicesPtrOfSlicesPtr(&slices)
}
