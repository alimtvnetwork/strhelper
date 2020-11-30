package strs

import "strings"

// Returns lower case strings by creating new pointer strings array.
//  (Don't modify in place, thus requires more memory consumption)
//
// Warning:
//  - Requires new memory space to do the case conversion and new spaces for the size of the array
//
// Invalid case (returns nil)
//  - if inputs == nil
func ToLowerPtrStrings(inputs *[]*string) *[]*string {
	if inputs == nil {
		return nil
	}

	newStrings := make([]*string, len(*inputs))

	for index, str := range *inputs {
		lowerStr := strings.ToLower(*str)
		newStrings[index] = &lowerStr
	}

	return &newStrings
}
