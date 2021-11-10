package isstrinternal

// EqualsPtr
//
// Returns :
//  - true : if both nil.
//  - false : if one nil and other not.
//  - true : if both are equal based on case sensitivity.
func EqualsPtr(first, second *string) bool {
	if first == nil && second == nil {
		return true
	}

	if first == nil && second != nil {
		return false
	}

	if first != nil && second == nil {
		return false
	}

	return first == second || *first == *second
}
