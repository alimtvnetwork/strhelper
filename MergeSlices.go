package strhelper

func MergeSlices(slices ...*[]string) []string {
	sliceLength := len(slices)

	if sliceLength == 0 {
		return []string{}
	}

	return *MergeSlicesPtrOfSlices(slices...)
}
