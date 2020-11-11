package strhelpercore

func convertPtrStringArrayToStringArray(ptrStrArray *[]*string) *[]string {
	newArray := make([]string, len(*ptrStrArray))

	for i, value := range *ptrStrArray {
		newArray[i] = *value
	}

	return &newArray
}
