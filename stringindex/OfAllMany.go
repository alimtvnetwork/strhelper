package stringindex

import (
	"gitlab.com/evatix-go/strhelper/strhelpercore"
)

// Find all the indexes for all the finding strings given.
//
// limits:
//  - How many indexes should we search for and then stop looking further.
//  - `-1` means find all, 0 => nil
func OfAllMany(
	content string,
	searchItems *[]string,
	startsAt int,
	limits int,
	isCaseSensitive bool,
) *strhelpercore.IndexesResultSet {
	searchRequestsMap := createDefaultSearchRequestsMap(
		searchItems,
		startsAt,
		limits,
		isCaseSensitive)

	return OfAllManyMapPtr(
		content,
		searchRequestsMap)
}
