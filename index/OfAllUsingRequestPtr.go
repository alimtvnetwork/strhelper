package index

import (
	"gitlab.com/evatix-go/core/constants"

	"gitlab.com/evatix-go/strhelper/internal/panichelper"
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

	wholeTextLength := len(*content)

	if request.StartsAt <= constants.InvalidNotFoundCase || request.StartsAt > wholeTextLength-1 {
		panichelper.StartAtIndexFailed(request.StartsAt, wholeTextLength)
	}

	if wholeTextLength > 0 && request.Search == "" {
		return getAllIndexesFromTheStartIndexGiven(
			wholeTextLength,
			request.StartsAt,
			request.Limits)
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

	// keep the default as best so that doesn't resize.
	defaultCapacity := wholeTextLength

	if wholeTextLength > constants.ArbitraryCapacity1000 {
		defaultCapacity = constants.ArbitraryCapacity100
	}

	if wholeTextLength > constants.ArbitraryCapacity250 {
		defaultCapacity = wholeTextLength / constants.ArbitraryCapacity10
	}

	indexes := make(
		[]int,
		constants.Zero,
		defaultCapacity)
	lastIndex := wholeTextLength - 1

	sendingRequest := strhelpercore.SearchRequest{
		Search:          *sendingSearchTerm,
		StartsAt:        request.StartsAt,
		Limits:          request.Limits,
		IsCaseSensitive: true,
	}

	searchIndividualRequest := strhelpercore.SearchIndividualRequest{
		Text:            sendingContent,
		SearchRequest:   &sendingRequest,
		WholeTextLength: wholeTextLength,
	}

	foundIndex := OfUsingRequestPtr(&searchIndividualRequest)

	if foundIndex > constants.InvalidNotFoundCase {
		indexes = append(indexes, foundIndex)
	}

	var nextIndex int

	for foundIndex > constants.InvalidNotFoundCase {
		nextIndex = foundIndex + 1
		if nextIndex > lastIndex || (request.Limits > -1 && len(indexes) >= request.Limits) {
			break
		}

		sendingRequest.StartsAt = nextIndex
		foundIndex = OfUsingRequestPtr(&searchIndividualRequest)

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

	return &indexes
}
