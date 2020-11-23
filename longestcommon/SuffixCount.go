package longestcommon

// Results count of suffix character matches. Where a, b can be at different lengths,
// it will find the longest common suffix.
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
//
// Recommendation : For performance use the [`...Ptr`] version of the method.
func SuffixCount(
	a, b string,
	bothLastIndexReduceBy int,
	isCaseSensitive bool,
) int {
	return SuffixCountPtr(
		&a,
		&b,
		bothLastIndexReduceBy,
		isCaseSensitive)
}
