package strhelper

import (
	"gitlab.com/evatix-go/strhelper/strhelpercore"
)

func ReplaceMultiple(
	text string,
	searchReplaceMap map[string]string,
	startsAt int,
	howManyReplace int,
	isCaseSensitive bool,
) string {
	return ReplaceMultiplePtr(
		&text,
		&searchReplaceMap,
		startsAt,
		howManyReplace,
		isCaseSensitive)
}

func ReplaceMultiplePtr(
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

	return ReplaceMultipleUsingRequestPtr(
		text,
		&searchReplaceRequestMap)
}

func ReplaceMultipleUsingRequestPtr(
	text *string,
	searchReplaceMap *map[string]strhelpercore.ReplaceIndividualRequest,
) string {
	request := strhelpercore.ReplaceRequestMultiple{
		Text:             text,
		SearchReplaceMap: searchReplaceMap,
	}

	return replaceMultipleInternalPtr(&request)
}
