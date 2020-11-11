package strhelper

import "strings"

// Assumptions here are a,b are not nil, at least empty string.
//
// Results count of prefix character matches. Where a, b can be at different lengths, it will find the longest common prefix(thus start matching).
//
// Returns
//
//  - count number of characters how much matches as prefix from (a,b).
//  - if no prefix found or both are empty string then returns 0 count value
//
// bothStartsAtIndex:
//  - `2` meaning both a,b start index comparing from this index 2, represents 3rd index.
//  - `0` means starts from 0
//
// Code Copied from Reference: https://bit.ly/35ZGJHc
func longestCommonPrefixCountInsensitive(
	a *string,
	b *string,
	bothStartsAtIndex int,
) int {
	lenA := len(*a)
	lenB := len(*b)

	al := strings.ToLower(*a)
	bl := strings.ToLower(*b)

	incrementing := 0
	for ; bothStartsAtIndex < lenA && bothStartsAtIndex < lenB; bothStartsAtIndex++ {
		if al[bothStartsAtIndex] != bl[bothStartsAtIndex] {
			return incrementing
		}

		incrementing++
	}

	return incrementing
}
