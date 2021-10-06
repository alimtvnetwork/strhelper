package index

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"

	"gitlab.com/evatix-go/strhelper/internal/indexinternal"
	"gitlab.com/evatix-go/strhelper/internal/panichelper"
)

// Returns all indexes where findingString is found.
//
// @limits:
//  - How many indexes should we search for and then stop looking further.
//  - `-1` means find all, 0 => nil
//
// Results:
//  - Invalid result can be nil if any (content == nil || findingString == nil) results nil.
//  - If no indexes found returns nil.
func OfLastAllPtr(
	content,
	findingString string,
	contentLengthDecreasedBy int,
	limits int,
	isCaseSensitive bool,
) []int {
	if limits == 0 {
		return nil
	}

	if content == findingString && contentLengthDecreasedBy == 0 {
		return []int{0}
	}

	wholeTextLength := len(content)
	searchLength := len(findingString)

	if searchLength > wholeTextLength-contentLengthDecreasedBy {
		return nil
	}

	if contentLengthDecreasedBy <= constants.InvalidNotFoundCase || contentLengthDecreasedBy > wholeTextLength-1 {
		panichelper.ContentLengthDecreasedByFailed(
			contentLengthDecreasedBy,
			wholeTextLength)
	}

	if wholeTextLength > 0 && searchLength == 0 {
		return getAllIndexesFromLastAndIndexGiven(
			wholeTextLength,
			contentLengthDecreasedBy,
			limits)
	}

	if limits == 0 {
		return nil
	}

	// making a copy of pointer only, not the object. copy of reference address
	// reference : https://play.golang.org/p/r65MrCg86YH
	sendingContent := content
	sendingSearchTerm := findingString

	if isCaseSensitive == false {
		// insensitive
		sendingContent = strings.ToLower(sendingContent)
		sendingSearchTerm = strings.ToLower(findingString)
	}

	// keep the default as best so that doesn't resize.
	return indexinternal.OfLastAllCaseSensitiveUsingLengthPtr(
		sendingContent,
		sendingSearchTerm,
		contentLengthDecreasedBy,
		limits,
		wholeTextLength,
		searchLength,
	)
}
