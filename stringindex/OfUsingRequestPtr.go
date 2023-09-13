package stringindex

import (
	"strings"

	"gitlab.com/auk-go/strhelper/internal/panichelper"
	"gitlab.com/auk-go/strhelper/strhelpercore"
)

// Returns the first index of the findingString in s
//
// Returns Index
//  - If text is found and nothing is invalid like (none is nil)
//
// Returns -1
//  - When not found or invalid case.
//
// Conditions (for panic):
//  - SearchRequest, SearchRequest.Text, SearchRequest.Search should NOT be nil.
//  - startsAt cannot be negative or greater than the length of text(s)
//goland:noinspection ALL
func OfUsingRequestPtr(searchIndividualRequest *strhelpercore.SearchIndividualRequest) int {
	if searchIndividualRequest == nil {
		panichelper.NullReferences(
			"searchIndividualRequest",
			"searchIndividualRequest.Text",
			"searchIndividualRequest.SearchRequest")
	}

	wholeTextLength := searchIndividualRequest.WholeTextLength
	searchRequest := searchIndividualRequest.SearchRequest

	if searchRequest.StartsAt < 0 || wholeTextLength-1 < searchRequest.StartsAt {
		panichelper.StartAtIndexFailed(
			searchRequest.StartsAt,
			wholeTextLength)
	}

	if searchRequest.IsCaseSensitive && searchRequest.StartsAt == 0 {
		return strings.Index(
			searchIndividualRequest.Text,
			searchRequest.Search)
	}

	if searchRequest.IsCaseSensitive {
		return ofCaseSensitiveUsingLengthPtr(
			searchIndividualRequest.Text,
			searchRequest.Search,
			searchRequest.StartsAt,
			wholeTextLength,
			len(searchRequest.Search))
	}

	return OfCaseInsensitive(
		searchIndividualRequest.Text,
		searchRequest.Search,
		searchRequest.StartsAt)
}
