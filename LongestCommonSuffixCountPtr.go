package strhelper

import "gitlab.com/evatix-go/strhelper/constants"

// Results count of suffix character matches. Where a, b can be at different lengths, it will find the longest common suffix.
//
// Returns
//
//  - count number of characters how much matches as suffix from (a,b).
//  - if no suffix found or both are empty string then returns 0 count value
//
// Conditions (Panic):
//  - if any (a,b) nil then panics
//  - bothLastIndexReduceBy cannot be negative
//
// bothLastIndexReduceBy:
//  - `2` meaning both a,b length consider len(a)-2, len(b)-2
//  - bothLastIndexReduceBy cannot be negative
//
// Code Copied from Reference: https://bit.ly/35ZGJHc
func LongestCommonSuffixCountPtr(
	a, b *string,
	bothLastIndexReduceBy int,
	isCaseSensitive bool,
) int {
	if a == nil || b == nil {
		panic("Either a or b is nil, please provide valid input at least empty string.")
	}

	if bothLastIndexReduceBy < 0 {
		panic("bothLastIndexReduceBy cannot be negative.")
	}

	lenA := len(*a)
	lenB := len(*b)

	if lenA == constants.Zero && ((lenB == constants.Zero && bothLastIndexReduceBy == constants.Zero) || lenB-1 >= bothLastIndexReduceBy) {
		return 0
	}

	if lenB == constants.Zero && lenA == constants.Zero && bothLastIndexReduceBy == constants.Zero {
		return 0
	}

	if !isCaseSensitive {
		// both needs to be in same case
		return longestSuffixCountInsensitive(
			a,
			b,
			bothLastIndexReduceBy)
	}

	incrementing := 0
	for ; bothLastIndexReduceBy < lenA && bothLastIndexReduceBy < lenB; bothLastIndexReduceBy++ {
		if (*a)[lenA-1-bothLastIndexReduceBy] != (*b)[lenB-1-bothLastIndexReduceBy] {
			return incrementing
		}

		incrementing++
	}

	return incrementing
}
