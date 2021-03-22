package isstrinternal

import "strings"

// Returns :
//  - true : if both nil.
//  - false : if one nil and other not.
//  - true : if both are equal based on case sensitivity.
//goland:noinspection ALL
func EqualsCasePtr(first, second *string, isCaseSensitive bool) bool {
	if first == nil && second == nil {
		return true
	}

	if first == nil && second != nil {
		return false
	}

	if first != nil && second == nil {
		return false
	}

	isEqualWithoutCase :=
		first == second ||
			*first == *second

	if isEqualWithoutCase || isCaseSensitive {
		// regardless true
		return isEqualWithoutCase
	}

	return strings.EqualFold(*first, *second)
}
