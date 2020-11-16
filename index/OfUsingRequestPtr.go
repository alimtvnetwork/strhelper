package index

import (
	"strings"

	"gitlab.com/evatix-go/strhelper/internal/pkg/constants"
	"gitlab.com/evatix-go/strhelper/internal/pkg/panichelper"
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
func OfUsingRequestPtr(searchIndividualRequest *strhelpercore.SearchIndividualRequest) int {
	if searchIndividualRequest == nil || searchIndividualRequest.Text == nil || searchIndividualRequest.SearchRequest == nil {
		panic(constants.SearchNullPanicMessage)
	}

	length := len(*searchIndividualRequest.Text)
	searchRequest := searchIndividualRequest.SearchRequest

	if searchRequest.StartsAt < 0 || length-1 < searchRequest.StartsAt {
		panichelper.StartAtIndexFailed(searchRequest.StartsAt, length)
	}

	if searchRequest.IsCaseSensitive && searchRequest.StartsAt == 0 {
		return strings.Index(*searchIndividualRequest.Text, searchRequest.Search)
	}

	if (*searchRequest).IsCaseSensitive {
		return OfCaseSensitive(
			searchIndividualRequest.Text,
			&searchRequest.Search,
			searchRequest.StartsAt)
	}

	return OfCaseInsensitive(
		searchIndividualRequest.Text,
		&searchRequest.Search,
		searchRequest.StartsAt)
}
