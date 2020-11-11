package strhelper

import (
	"gitlab.com/evatix-go/strhelper/strhelpercore"
)

func MultiStrIndexesOfAllUsingMapPtr(
	content *string,
	searchRequestsMap *map[string]strhelpercore.SearchRequest,
) strhelpercore.IndexesResultSet {
	if content == nil || searchRequestsMap == nil {
		return strhelpercore.NewEmptyIndexesResultSet()
	}

	searchingItemsLength := len(*searchRequestsMap)

	if searchingItemsLength == 0 {
		return strhelpercore.NewEmptyIndexesResultSet()
	}

	indexesMap := make(map[string][]int, searchingItemsLength)
	hasFoundAny := false

	for key, searchRequest := range *searchRequestsMap {
		indexes := IndexesOfAllPtr(
			content,
			&key,
			searchRequest.StartsAt,
			searchRequest.IsCaseSensitive)

		if indexes != nil && len(indexes) > 0 {
			indexesMap[key] = indexes
			hasFoundAny = true
		} else {
			indexesMap[key] = nil
		}
	}

	return strhelpercore.NewIndexesResultSet(
		&indexesMap,
		hasFoundAny)
}
