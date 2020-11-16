package longestcommon

import (
	"gitlab.com/evatix-go/strhelper/strconst"
)

// Assumptions here are a,b are not nil, at least empty string.
//
// Results count of prefix character matches. Where a, b can be at different lengths, it will find the longest common prefix(thus start matching).
//
// Returns
//
//  - count number of characters how much matches as prefix from (a,b).
//  - if no prefix found or both are empty string then returns 0 count value
//
// Conditions (Panic):
//  - if any (a,b) nil then panics
//  - if bothStartsAtIndex negative
//
// bothStartsAtIndex:
//  - `2` meaning both a,b start index comparing from this index 2, represents 3rd index.
//  - if less than 0 then panic
//
// Code Copied from Reference: https://bit.ly/35ZGJHc
func PrefixCountPtr(
	a, b *string,
	bothStartAtIndex int,
	isCaseSensitive bool,
) int {
	if a == nil || b == nil {
		panic("Either a or b is nil, please provide valid input at least empty string.")
	}

	if bothStartAtIndex < 0 {
		panic("bothStartAtIndex cannot be negative.")
	}

	lenA := len(*a)
	lenB := len(*b)

	if lenA == strconst.Zero && ((lenB == strconst.Zero && bothStartAtIndex == strconst.Zero) || lenB-1 >= bothStartAtIndex) {
		return 0
	}

	if lenB == strconst.Zero && lenA == strconst.Zero && bothStartAtIndex == strconst.Zero {
		return 0
	}

	if !isCaseSensitive {
		// both needs to be in same case
		return longestCommonPrefixCountInsensitive(
			a,
			b,
			bothStartAtIndex)
	}

	return longestCommonPrefixSensitiveCase(
		a,
		b,
		bothStartAtIndex)
}
