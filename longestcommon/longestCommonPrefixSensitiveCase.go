package longestcommon

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
//
// Code Copied from Reference: https://bit.ly/35ZGJHc
func longestCommonPrefixSensitiveCase(
	a *string,
	b *string,
	bothStartAtIndex int,
) int {
	lenA := len(*a)
	lenB := len(*b)

	incrementing := 0
	for ; bothStartAtIndex < lenA && bothStartAtIndex < lenB; bothStartAtIndex++ {
		if (*a)[bothStartAtIndex] != (*b)[bothStartAtIndex] {
			return incrementing
		}

		incrementing++
	}

	return incrementing
}
