package chars

// CountCharSensitive
//
// Counts and returns the count number based on chars present in the str (case: Sensitive)
// Returns 0 if str is nil or empty string.
func CountCharSensitive(str *string, char1 uint8, startsAt int) int {
	length := len(*str)
	found := 0

	if length == 0 {
		return found
	}

	for ; startsAt < length; startsAt++ {
		char := (*str)[startsAt]
		if char == char1 {
			found++
		}
	}

	return found
}
