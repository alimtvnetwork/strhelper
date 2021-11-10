package isstr

import "strings"

// Equals
//
// Returns :
//  - true : if both are equal based on case sensitivity.
func Equals(first, second string, isCaseSensitive bool) bool {
	if isCaseSensitive {
		return first == second
	}

	// insensitive
	return strings.EqualFold(first, second)
}
