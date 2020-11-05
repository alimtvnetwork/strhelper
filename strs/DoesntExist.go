package strs

// Returns true if the findingString NOT exist in the array, if array is empty or nil then returns true.
func DoesntExist(lines *[]string, findingString *string, isCaseSensitive bool) bool {
	return !IsExists(lines, findingString, isCaseSensitive)
}
