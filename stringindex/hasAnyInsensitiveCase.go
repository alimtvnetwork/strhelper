package stringindex

import "gitlab.com/auk-go/strhelper/strhelpercore"

func hasAnyInsensitiveCase(searchRequestsMap map[string]strhelpercore.SearchRequest) bool {
	for _, searchRequest := range searchRequestsMap {
		if searchRequest.IsCaseSensitive == false {
			// insensitive
			return true
		}
	}

	return false
}
