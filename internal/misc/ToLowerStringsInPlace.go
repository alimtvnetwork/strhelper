package misc

import "strings"

// Returns all the strings array to lower case form by modifying strings array.
//
// Even-though content is modified in place, best to work with the return value.
//
// Warning:
//  - Modifies content in place.
//  - Still it requires new memory space to do the case conversion because string is immutable.
//
// Invalid case (returns nil)
//  - if inputs == nil
func ToLowerStringsInPlace(inputs *[]string) *[]string {
	if inputs == nil {
		return nil
	}

	for index, str := range *inputs {
		(*inputs)[index] = strings.ToLower(str)
	}

	return inputs
}
