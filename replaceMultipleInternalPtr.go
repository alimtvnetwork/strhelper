package strhelper

import (
	"gitlab.com/evatix-go/strhelper/strhelpercore"
)

func replaceMultipleInternalPtr(request *strhelpercore.ReplaceRequestMultiple) string {
	if len(*request.SearchReplaceMap) == 0 {
		return *request.Text
	}

	if *request.Text == "" {
		return emptyReplaceResultUsingRequestMultiple(request)
	}

	searchMap := convertSearchReplaceMapToSearchMap(request.SearchReplaceMap)

	indexesResultSet := MultiStrIndexesOfAllUsingMapPtr(
		request.Text,
		searchMap)

	if indexesResultSet.IsEmpty() {
		// returns as is
		return *request.Text
	}

	textLength := len(*request.Text)
	replaceCount := indexesResultSet.CountOfAllFoundIndexes()

	// not found case
	if replaceCount == 0 {
		return *request.Text
	}

	// all length calculation
	changeInLengthNewString := getChangeInNewLengthAndReplaceCountUpdate(request, indexesResultSet)

	// Apply replacements to buffer.
	indexesAsKeyMap := indexesResultSet.GetIndexesMapWhereIndexAsKey()

	return getCompiledReplaceMultiple(
		request,
		textLength,
		changeInLengthNewString,
		indexesAsKeyMap)
}

func getCompiledReplaceMultiple(
	request *strhelpercore.ReplaceRequestMultiple,
	textLength int,
	changeInLengthNewString int,
	indexesAsKeyMap *map[int]string,
) string {
	chars := make([]byte, textLength+changeInLengthNewString)
	wordIndex := 0
	for i := 0; i < textLength; i++ {
		searchStr, isIndexExist := (*indexesAsKeyMap)[i]

		if isIndexExist {
			replaceRequest := (*request.SearchReplaceMap)[searchStr]

			if replaceRequest.ShouldReplace() {
				// too many nesting
				// keeping as is for performance
				// found modify, reference : https://bit.ly/2IAEJNe
				wordIndex += copy(chars[wordIndex:], replaceRequest.ReplaceWith)
				i += len(searchStr) - 1 // we should skip the search text since already replaced.
				replaceRequest.ReplaceCountDecrease()
				(*request.SearchReplaceMap)[searchStr] = replaceRequest
				continue
			}
		}

		// not found existing, keep as is
		chars[wordIndex] = (*(*request).Text)[i]
		wordIndex++
	}

	return string(chars)
}

// Updates replace count on indexesResultSet,
//
// and returns the final length for the new replaced text.
func getChangeInNewLengthAndReplaceCountUpdate(
	request *strhelpercore.ReplaceRequestMultiple,
	indexesResultSet *strhelpercore.IndexesResultSet,
) int {
	changeInLengthNewString := 0

	for searchKey, indexes := range *indexesResultSet.StringKeyAsIndexesMap {
		replaceIndividualRequest := (*request.SearchReplaceMap)[searchKey]
		replaceStr := replaceIndividualRequest.ReplaceWith
		replaceCount := len(*indexes)
		replaceIndividualRequest.SetReplaceCount(replaceCount)
		(*request.SearchReplaceMap)[searchKey] = replaceIndividualRequest
		changeInLengthNewString += replaceCount * (len(replaceStr) - len(searchKey))
	}

	return changeInLengthNewString
}

func emptyReplaceResultUsingRequestMultiple(request *strhelpercore.ReplaceRequestMultiple) string {
	result, hasAny := (*request.SearchReplaceMap)[*request.Text]

	if hasAny {
		return result.ReplaceWith
	}

	return *request.Text
}
