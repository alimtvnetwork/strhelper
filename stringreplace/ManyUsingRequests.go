package stringreplace

import "gitlab.com/auk-go/strhelper/strhelpercore"

func ManyUsingRequests(
	text string,
	searchReplaceMap map[string]strhelpercore.ReplaceIndividualRequest,
) string {
	request := strhelpercore.ReplaceRequestMultiple{
		Text:             text,
		SearchReplaceMap: searchReplaceMap,
	}

	return replaceMultipleInternalPtr(&request)
}
