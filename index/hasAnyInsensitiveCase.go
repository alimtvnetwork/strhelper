package index

import "gitlab.com/evatix-go/strhelper/strhelpercore"

func hasAnyInsensitiveCase(searchRequestsMap *map[string]strhelpercore.SearchRequest) bool {
	for _, searchRequest := range *searchRequestsMap {
		if searchRequest.IsCaseSensitive == false {
			// insensitive
			return true
		}
	}

	return false
}
