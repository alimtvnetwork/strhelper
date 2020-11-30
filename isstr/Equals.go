package isstr

import "strings"

// Returns :
//  - true : if both are equal based on case sensitivity.
func Equals(first, second string, isCaseSensitive bool) bool {
	if isCaseSensitive {
		return first == second
	}

	// insensitive
	fLower := strings.ToLower(first)
	sLower := strings.ToLower(second)

	return fLower == sLower
}
