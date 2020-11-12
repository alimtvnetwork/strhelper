package strs

import "strings"

// Returns upper case strings by creating new pointer strings array.
//  (Don't modify in place, thus requires more memory consumption)
//
// Warning:
//  - Requires new memory space to do the case conversion and new spaces for the size of the array
//
// Invalid case (returns nil)
//  - if inputs == nil
func ToUpperPtrStrings(inputs *[]*string) *[]*string {
	if inputs == nil {
		return nil
	}

	newStrings := make([]*string, len(*inputs))

	for index, str := range *inputs {
		upperStr := strings.ToUpper(*str)
		newStrings[index] = &upperStr
	}

	return &newStrings
}
