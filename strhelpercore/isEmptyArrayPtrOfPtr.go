package strhelpercore

func isEmptyStringArrayPtr(strPointers *[]string) bool {
	return strPointers == nil || len(*strPointers) == 0
}

func isEmptyIntArrayOfArrayPtr(intPointers *[][]int) bool {
	return intPointers == nil || len(*intPointers) == 0
}
