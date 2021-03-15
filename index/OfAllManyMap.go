package index

import "gitlab.com/evatix-go/strhelper/strhelpercore"

// Find all the indexes for all the finding strings given.
//
// searchMap:
//   - Key : What to search for
//   - Value : Where starts at, usually 0 for default start.
// limits:
//  - How many indexes should we search for and then stop looking further.
//  - `-1` means find all, 0 => nil
func OfAllManyMap(
	content *string,
	searchMap *map[string]int,
	limits int,
	isCaseSensitive bool,
) *strhelpercore.IndexesResultSet {
	searchRequestsMap := createSearchRequestsMap(
		searchMap,
		limits,
		isCaseSensitive)

	return OfAllManyMapPtr(
		content,
		searchRequestsMap)
}
