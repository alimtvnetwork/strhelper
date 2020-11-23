package converters

// StringArrayToPtrStringArray will give empty or converted results array (not nil)
func StringArrayToPtrStringArray(ptrStrArray *[]string) *[]*string {
	if ptrStrArray == nil || *ptrStrArray == nil {
		var emptyResult []*string

		return &emptyResult
	}

	newArray := make([]*string, len(*ptrStrArray))

	for i, value := range *ptrStrArray {
		newArray[i] = &value
	}

	return &newArray
}
