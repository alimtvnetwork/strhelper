package strhelpercore

func isEmptyArrayPtrOfPtr(interfaces *[]*interface{}) bool {
	return interfaces == nil || len(*interfaces) == 0
}

func isEmptyStringArrayPtr(strPointers *[]string) bool {
	return strPointers == nil || len(*strPointers) == 0
}

func isEmptyStringArrayPtrOfPtr(strPointers *[]*string) bool {
	return strPointers == nil || len(*strPointers) == 0
}

func isEmptyIntArrayOfArrayPtr(intPointers *[][]int) bool {
	return intPointers == nil || len(*intPointers) == 0
}
