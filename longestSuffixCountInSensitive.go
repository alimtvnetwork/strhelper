package strhelper

import "strings"

// Assumptions here are a,b are not nil, at least empty string.
//
// Results count of suffix character matches. Where a, b can be at different lengths, it will find the longest common suffix.
//
// Returns
//
//  - count number of characters how much matches as suffix from (a,b).
//  - if no suffix found or both are empty string then returns 0 count value
//
// Conditions (Panic):
//  - if any (a,b) nil then panics
//
// bothLastIndexReduceBy:
//  - `2` meaning both a,b length consider len(a)-2, len(b)-2
//
// Code Copied from Reference: https://bit.ly/35ZGJHc
func longestSuffixCountInsensitive(
	a *string,
	b *string,
	bothLastIndexReduceBy int,
) int {
	lenA := len(*a)
	lenB := len(*b)

	al := strings.ToLower(*a)
	bl := strings.ToLower(*b)

	incrementing := 0

	for ; bothLastIndexReduceBy < lenA && bothLastIndexReduceBy < lenB; bothLastIndexReduceBy++ {
		if al[lenA-1-bothLastIndexReduceBy] != bl[lenB-1-bothLastIndexReduceBy] {
			return incrementing
		}

		incrementing++
	}

	return incrementing
}
