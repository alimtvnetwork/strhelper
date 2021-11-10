package isstr

import "strings"

// EqualsPtr
//
// Returns :
//  - true : if both nil.
//  - false : if one nil and other not.
//  - true : if both are equal based on case sensitivity.
func EqualsPtr(first, second *string, isCaseSensitive bool) bool {
	if first == nil && second == nil {
		return true
	}

	if first == nil && second != nil {
		return false
	}

	if first != nil && second == nil {
		return false
	}

	if isCaseSensitive {
		return first == second || *first == *second
	}

	// insensitive
	return strings.EqualFold(*first, *second)
}
