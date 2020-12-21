package index

import (
	"gitlab.com/evatix-go/strhelper/internal/panichelper"
	"gitlab.com/evatix-go/strhelper/strconst"
	"gitlab.com/evatix-go/strhelper/strhelpercore"
	"gitlab.com/evatix-go/strhelper/strto"
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
func OfAllUsingRequestPtr(
	content *string,
	request *strhelpercore.SearchRequest,
) *[]int {
	if content == nil || request == nil {
		return nil
	}

	length := len(*content)

	if request.StartsAt <= strconst.InvalidNotFoundCase || request.StartsAt > length-1 {
		panichelper.StartAtIndexFailed(request.StartsAt, length)
	}

	if length > 0 && request.Search == "" {
		return getAllIndexesFromTheStartIndexGiven(length, request.StartsAt)
	}

	// making a copy of pointer only, not the object. copy of reference address
	// reference : https://play.golang.org/p/r65MrCg86YH
	sendingContent := content
	sendingSearchTerm := &request.Search

	if request.IsCaseSensitive == false {
		// insensitive
		sendingContent = strto.LowerStrPtr(sendingContent)
		sendingSearchTerm = strto.LowerStrPtr(sendingSearchTerm)
	}

	indexes := make([]int, strconst.Zero, length)
	lastIndex := length - 1

	sendingRequest := strhelpercore.SearchRequest{
		Search:          *sendingSearchTerm,
		StartsAt:        request.StartsAt,
		Limits:          request.Limits,
		IsCaseSensitive: true,
	}

	searchIndividualRequest := strhelpercore.SearchIndividualRequest{
		Text:          sendingContent,
		SearchRequest: &sendingRequest,
	}

	foundIndex := OfUsingRequestPtr(&searchIndividualRequest)

	if foundIndex > strconst.InvalidNotFoundCase {
		indexes = append(indexes, foundIndex)
	}

	var nextIndex int

	for foundIndex > strconst.InvalidNotFoundCase {
		nextIndex = foundIndex + 1
		if nextIndex > lastIndex || (request.Limits > -1 && len(indexes) >= request.Limits) {
			break
		}

		sendingRequest.StartsAt = nextIndex
		foundIndex = OfUsingRequestPtr(&searchIndividualRequest)

		if foundIndex > strconst.InvalidNotFoundCase {
			indexes = append(indexes, foundIndex)
		} else {
			// not found at any, will not continue
			break
		}
	}

	if len(indexes) == strconst.Zero {
		return nil
	}

	return &indexes
}
