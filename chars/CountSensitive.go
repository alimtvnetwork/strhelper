package chars

// Counts and returns the count number based on chars present in the str (case: Sensitive)
// Returns 0 if str is nil or empty string.
func CountCharSensitive(str *string, char1 uint8, startAt int) int {
	length := len(*str)
	found := 0

	if length == 0 {
		return found
	}

	for ; startAt < length; startAt++ {
		char := (*str)[startAt]
		if char == char1 {
			found++
		}
	}

	return found
}
