package strhelper

import (
	"strings"

	"gitlab.com/evatix-go/strhelper/strhelpercore"
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
func IndexOfUsingRequestPtr(searchIndividualRequest *strhelpercore.SearchIndividualRequest) int {
	if searchIndividualRequest == nil || searchIndividualRequest.Text == nil || searchIndividualRequest.SearchRequest == nil {
		panic(searchNullPanicMessage)
	}

	length := len(*searchIndividualRequest.Text)
	searchRequest := searchIndividualRequest.SearchRequest

	if searchRequest.StartsAt < 0 || length-1 < searchRequest.StartsAt {
		startAtIndexFailed(searchRequest.StartsAt, length)
	}

	if searchRequest.IsCaseSensitive && searchRequest.StartsAt == 0 {
		return strings.Index(*searchIndividualRequest.Text, searchRequest.Search)
	}

	if (*searchRequest).IsCaseSensitive {
		return IndexOfCaseSensitive(
			searchIndividualRequest.Text,
			&searchRequest.Search,
			searchRequest.StartsAt)
	}

	return IndexOfCaseInsensitive(
		searchIndividualRequest.Text,
		&searchRequest.Search,
		searchRequest.StartsAt)
}
