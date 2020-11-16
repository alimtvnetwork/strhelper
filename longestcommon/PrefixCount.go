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
// Conditions (Panic):
//  - if any (a,b) nil then panics
//  - if bothStartsAtIndex negative
//
// bothStartsAtIndex:
//  - `2` meaning both a,b start index comparing from this index 2, represents 3rd index.
//  - if less than 0 then panic
//
// Code Copied from Reference: https://bit.ly/35ZGJHc
//
// Recommendation : For performance use the [`...Ptr`] version of the method.
func PrefixCount(
	a, b string,
	bothStartsAtIndex int,
	isCaseSensitive bool,
) int {
	return PrefixCountPtr(
		&a,
		&b,
		bothStartsAtIndex,
		isCaseSensitive)
}
