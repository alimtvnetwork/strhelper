package stringreplace

import "gitlab.com/evatix-go/strhelper/strhelpercore"

func ManyPtr(
	text *string,
	searchReplaceMap *map[string]string,
	startsAt int,
	howManyReplace int,
	isCaseSensitive bool,
) string {
	if searchReplaceMap == nil || len(*searchReplaceMap) == 0 {
		return *text
	}

	searchReplaceRequestMap := map[string]strhelpercore.ReplaceIndividualRequest{}

	for key, value := range *searchReplaceMap {
		searchReplaceRequestMap[key] = strhelpercore.ReplaceIndividualRequest{
			Search:          key,
			ReplaceWith:     value,
			StartsAt:        startsAt,
			HowManyReplace:  howManyReplace,
			IsCaseSensitive: isCaseSensitive,
		}
	}

	return ManyUsingRequests(
		text,
		&searchReplaceRequestMap)
}
