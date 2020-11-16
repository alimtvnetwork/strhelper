package index

import (
	"gitlab.com/evatix-go/strhelper/strhelpercore"
)

// Find all the indexes for all the finding strings given.
//
// limits:
//  - How many indexes should we search for and then stop looking further.
//  - `-1` means find all
func OfAllMany(
	content *string,
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

// Find all the indexes for all the finding strings given.
//
// searchMap:
//   - Key : What to search for
//   - Value : Where starts at, usually 0 for default start.
// limits:
//  - How many indexes should we search for and then stop looking further.
//  - `-1` means find all
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

// Find all the indexes for all the finding strings given.
func OfAllManyMapPtr(
	content *string,
	searchRequestsMap *map[string]strhelpercore.SearchRequest,
) *strhelpercore.IndexesResultSet {
	if content == nil || searchRequestsMap == nil {
		return strhelpercore.NewEmptyIndexesResultSet()
	}

	searchingItemsLength := len(*searchRequestsMap)

	if searchingItemsLength == 0 {
		return strhelpercore.NewEmptyIndexesResultSet()
	}

	indexesMap := make(map[string]*[]int, searchingItemsLength)
	var hasFoundAny bool
	var totalLength, maxIndexFound int
	maxIndexFound = -1

	// making a copy of pointer only, not the object. copy of reference address
	// reference : https://play.golang.org/p/r65MrCg86YH
	var lowerCaseContent, sendingContent *string
	if hasAnyInsensitiveCase(searchRequestsMap) {
		lowerCaseContent = strto.LowerStrPtr(content)
	}

	for key, searchRequest := range *searchRequestsMap {
		sendingContent = content

		if searchRequest.IsCaseSensitive == false {
			// insensitive
			sendingContent = lowerCaseContent
			searchRequest.Search = *strto.LowerStrPtr(&searchRequest.Search)
			searchRequest.IsCaseSensitive = true
		}

		indexes := OfAllUsingRequestPtr(
			sendingContent,
			&searchRequest)

		if indexes == nil || *indexes == nil {
			indexesMap[key] = nil
			continue
		}

		totalLength = len(*indexes)
		indexesMap[key] = indexes
		hasFoundAny = true

		if maxIndexFound < (*indexes)[totalLength-1] {
			maxIndexFound = (*indexes)[totalLength-1]
		}
	}

	return strhelpercore.NewIndexesResultSet(
		&indexesMap,
		hasFoundAny,
		maxIndexFound)
}

func hasAnyInsensitiveCase(searchRequestsMap *map[string]strhelpercore.SearchRequest) bool {
	for _, searchRequest := range *searchRequestsMap {
		if searchRequest.IsCaseSensitive == false {
			// insensitive
			return true
		}
	}

	return false
}
