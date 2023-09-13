package stringindex

import "gitlab.com/auk-go/strhelper/strhelpercore"

// searchMap:
//   - Key : What to search for
//   - Value : Where starts at, usually 0 for default start.
// limits:
//  - How many indexes should we search for and then stop looking further.
//  - `-1` means find all
func createSearchRequestsMap(
	searchMap map[string]int,
	limits int,
	isCaseSensitive bool,
) map[string]strhelpercore.SearchRequest {
	if searchMap == nil || len(searchMap) == 0 {
		return nil
	}

	searchRequestsMap := make(map[string]strhelpercore.SearchRequest, len(searchMap))

	for key, startsAt := range searchMap {
		searchRequestsMap[key] = strhelpercore.SearchRequest{
			Search:          key,
			StartsAt:        startsAt,
			Limits:          limits,
			IsCaseSensitive: isCaseSensitive,
		}
	}

	return searchRequestsMap
}

// limits:
//  - How many indexes should we search for and then stop looking further.
//  - `-1` means find all
func createDefaultSearchRequestsMap(
	searchItems *[]string,
	startsAt int,
	limits int,
	isCaseSensitive bool,
) map[string]strhelpercore.SearchRequest {
	if searchItems == nil || len(*searchItems) == 0 {
		return nil
	}

	searchRequestsMap := make(map[string]strhelpercore.SearchRequest, len(*searchItems))

	for _, search := range *searchItems {
		searchRequestsMap[search] = strhelpercore.SearchRequest{
			Search:          search,
			StartsAt:        startsAt,
			Limits:          limits,
			IsCaseSensitive: isCaseSensitive,
		}
	}

	return searchRequestsMap
}
