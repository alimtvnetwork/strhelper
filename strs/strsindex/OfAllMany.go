package strsindex

import (
	"gitlab.com/evatix-go/strhelper/internal/pkg/isstrsinternal"
	"gitlab.com/evatix-go/strhelper/internal/pkg/panichelper"
	"gitlab.com/evatix-go/strhelper/strconst"
	"gitlab.com/evatix-go/strhelper/strhelpercore"
	"gitlab.com/evatix-go/strhelper/strs"
)

// Returns all indexes by searching all @findingMap items.
//
// @limits:
//  - How many indexes should we search for and then stop looking further.
//  - `-1` means find all
//
// Results:
//  - Invalid result can be nil if any (lines == nil || searchTerms == nil) results nil.
//  - If no indexes found returns nil.
func OfAllMany(
	lines *[]string,
	searchTerms *[]string,
	startsAtIndex int,
	limits int,
	isCaseSensitive bool,
) *strhelpercore.IndexesResultSet {
	if isstrsinternal.EmptyPtr(lines) || isstrsinternal.EmptyPtr(searchTerms) {
		return nil
	}

	searchingItemsLength := len(*searchTerms)

	if searchingItemsLength == 0 {
		return strhelpercore.NewEmptyIndexesResultSet()
	}

	length := len(*lines)

	if startsAtIndex <= strconst.InvalidNotFoundCase || startsAtIndex > length-1 {
		panichelper.StartAtIndexFailed(startsAtIndex, length)
	}

	// making a copy of pointer only, not the object. copy of reference address
	// reference : https://play.golang.org/p/r65MrCg86YH
	sendingLines := lines
	sendingSearchTerms := searchTerms

	if isCaseSensitive == false {
		// insensitive
		sendingLines = strs.ToLowerStrings(lines)
		sendingSearchTerms = strs.ToLowerStrings(searchTerms)
	}

	indexesMap := make(map[string]*[]int, searchingItemsLength)
	var hasFoundAny bool
	var totalLength, maxIndexFound int
	maxIndexFound = -1

	for _, searchTerm := range *sendingSearchTerms {
		indexes := OfAll(
			sendingLines,
			&searchTerm,
			startsAtIndex,
			limits,
			true)

		if indexes == nil || *indexes == nil {
			indexesMap[searchTerm] = nil
			continue
		}

		totalLength = len(*indexes)
		indexesMap[searchTerm] = indexes
		hasFoundAny = true

		if maxIndexFound < (*indexes)[totalLength-1] {
			maxIndexFound = (*indexes)[totalLength-1]
		}
	}

	return strhelpercore.NewIndexesResultSet(
		&indexesMap,
		hasFoundAny,
		maxIndexFound)
}
