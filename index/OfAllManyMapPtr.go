package index

import (
	"gitlab.com/evatix-go/strhelper/strhelpercore"
	"gitlab.com/evatix-go/strhelper/strto"
)

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
	var lastIndex, maxIndexFound, lastIndexValue int
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

		lastIndex = len(*indexes) - 1
		indexesMap[key] = indexes
		hasFoundAny = true
		lastIndexValue = (*indexes)[lastIndex]

		if maxIndexFound < lastIndexValue {
			maxIndexFound = lastIndexValue
		}
	}

	return strhelpercore.NewIndexesResultSet(
		&indexesMap,
		hasFoundAny,
		maxIndexFound)
}
