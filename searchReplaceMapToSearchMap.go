package strhelper

import (
	"gitlab.com/evatix-go/strhelper/strhelpercore"
)

func convertSearchReplaceMapToSearchMap(
	searchReplaceMap *map[string]strhelpercore.ReplaceIndividualRequest,
) *map[string]strhelpercore.SearchRequest {
	newMap := make(map[string]strhelpercore.SearchRequest, len(*searchReplaceMap))

	for key, replaceRequest := range * searchReplaceMap {
		newMap[key] = strhelpercore.SearchRequest{
			Search:          replaceRequest.Search,
			StartsAt:        replaceRequest.StartsAt,
			HowManyReplace:  replaceRequest.HowManyReplace,
			IsCaseSensitive: replaceRequest.IsCaseSensitive,
		}
	}

	return &newMap
}
