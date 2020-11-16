package longestcommon

import (
	"gitlab.com/evatix-go/strhelper/strconst"
)

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
func SuffixCountPtr(
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

	if lenA == strconst.Zero && ((lenB == strconst.Zero && bothLastIndexReduceBy == strconst.Zero) || lenB-1 >= bothLastIndexReduceBy) {
		return 0
	}

	if lenB == strconst.Zero && lenA == strconst.Zero && bothLastIndexReduceBy == strconst.Zero {
		return 0
	}

	if !isCaseSensitive {
		// both needs to be in same case
		return longestCommonSuffixCountInsensitive(
			a,
			b,
			bothLastIndexReduceBy)
	}

	return longestCommonSuffixCount(
		a,
		b,
		bothLastIndexReduceBy)
}
