package isstrs

// ContainsSimple Returns true if the findingString present in the array,
// if array is empty or nil then returns false.
//
// One can use Exists similar to contains has fewer arguments
func ContainsSimple(lines []string, line string) bool {
	if len(lines) == 0 {
		return false
	}

	for _, sliceItem := range lines {
		if line == sliceItem {
			return true
		}
	}

	return false
}
