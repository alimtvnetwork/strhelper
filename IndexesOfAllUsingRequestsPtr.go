package strhelper

import (
	"gitlab.com/evatix-go/strhelper/constants"
	"gitlab.com/evatix-go/strhelper/strhelpercore"
)

// Returns all indexes where findingString is found.
//
// @limits:
//  - How many indexes should we search for and then stop looking further.
//  - `-1` means find all
//
// Results:
//  - Invalid result can be nil if any (content == nil || findingString == nil) results nil.
//  - If no indexes found returns nil.
func IndexesOfAllUsingRequestPtr(
	content *string,
	request *strhelpercore.SearchRequest,
) []int {
	if content == nil || request == nil {
		return nil
	}

	length := len(*content)

	if request.StartsAt <= constants.InvalidNotFoundCase || request.StartsAt > length-1 {
		startAtIndexFailed(request.StartsAt, length)
	}

	if length > 0 && request.Search == "" {
		return getAllIndexesFromTheStartIndexGiven(length, request.StartsAt)
	}

	indexes := make([]int, constants.Zero, length)

	lastIndex := length - 1

	searchIndividualRequest := strhelpercore.SearchIndividualRequest{
		Text:          content,
		SearchRequest: request,
	}

	foundIndex := IndexOfUsingRequestPtr(&searchIndividualRequest)

	if foundIndex > constants.InvalidNotFoundCase {
		indexes = append(indexes, foundIndex)
	}

	request2 := strhelpercore.SearchRequest{
		Search:          request.Search,
		StartsAt:        foundIndex+1,
		Limits:          request.Limits,
		IsCaseSensitive: request.IsCaseSensitive,
	}

	searchIndividualRequest.SearchRequest = &request2
	var nextIndex int

	for foundIndex > constants.InvalidNotFoundCase {
		nextIndex = foundIndex + 1
		if nextIndex > lastIndex || (request.Limits > -1 && len(indexes) >= request.Limits) {
			break
		}

		request2.StartsAt = nextIndex
		foundIndex = IndexOfUsingRequestPtr(&searchIndividualRequest)

		if foundIndex > constants.InvalidNotFoundCase {
			indexes = append(indexes, foundIndex)
		} else {
			// not found at any, will not continue
			break
		}
	}

	if len(indexes) == constants.Zero {
		return nil
	}

	return indexes
}
