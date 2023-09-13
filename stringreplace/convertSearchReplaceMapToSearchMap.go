package stringreplace

import "gitlab.com/auk-go/strhelper/strhelpercore"

func convertSearchReplaceMapToSearchMap(
	searchReplaceMap map[string]strhelpercore.ReplaceIndividualRequest,
) map[string]strhelpercore.SearchRequest {
	newMap := make(map[string]strhelpercore.SearchRequest, len(searchReplaceMap))

	for key, replaceRequest := range searchReplaceMap {
		newMap[key] = strhelpercore.SearchRequest{
			Search:          replaceRequest.Search,
			StartsAt:        replaceRequest.StartsAt,
			Limits:          replaceRequest.HowManyReplace,
			IsCaseSensitive: replaceRequest.IsCaseSensitive,
		}
	}

	return newMap
}
