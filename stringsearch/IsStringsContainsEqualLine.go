package stringsearch

func IsStringsContainsEqualLine(slice []string, line string) bool {
	if len(slice) == 0 {
		return false
	}

	for _, sliceItem := range slice {
		if line == sliceItem {
			return true
		}
	}

	return false
}
