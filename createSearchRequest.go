package strhelper

import "gitlab.com/evatix-go/strhelper/strhelpercore"

// @limits:
//  - How many indexes should we search for and then stop looking further.
//  - `-1` means find all
func createSearchRequest(
	findingString *string,
	startsAtIndex int,
	limits int,
	isCaseSensitive bool,
) *strhelpercore.SearchRequest {
	return &strhelpercore.SearchRequest{
		Search:          *findingString,
		StartsAt:        startsAtIndex,
		Limits:          limits,
		IsCaseSensitive: isCaseSensitive,
	}
}
