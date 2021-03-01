package splits

import (
	"gitlab.com/evatix-go/core/constants"

	"gitlab.com/evatix-go/strhelper/index"
	"gitlab.com/evatix-go/strhelper/isstr"
	"gitlab.com/evatix-go/strhelper/strhelpercore"
	"gitlab.com/evatix-go/strhelper/whitespace"
)

// Multiple split occur from the given array of splits.
//
// Basics of split("Hello World", " ") -> ["Hello", "World"] splitter will not be available in the result.
//
// limit (Default):
//  - Given as `-1`
//
// splitStartsAt (default):
//  - Given as `0`
// isCaseSensitive (default):
//  - true
func ManyWithDefaults(
	str string,
	splitsBy ...string,
) *strhelpercore.SplitResultOverview {
	return ManyPtr(
		&str,
		&splitsBy,
		0,
		constants.InvalidNotFoundCase,
		true)
}

// Multiple split occur from the given array of splits.
//
// Basics of split("Hello World", " ") -> ["Hello", "World"] splitter will not be available in the result.
//
// limit :
//  - number of times split will performed for all
//  - if -1 then all split will occur
//
// splitStartsAt:
//  - where split searching will start from.
func Many(
	str string,
	splitStartsAt,
	limit int,
	isCaseSensitive bool,
	splitsBy ...string,
) *strhelpercore.SplitResultOverview {
	return ManyPtr(
		&str,
		&splitsBy,
		splitStartsAt,
		limit,
		isCaseSensitive)
}

// Multiple split occur from the given array of splits.
//
// Basics of split("Hello World", " ") -> ["Hello", "World"] splitter will not be available in the result.
//
// @limit (Default):
//  - Given as `-1`
// @splitStartsAt (default):
//  - Given as `0`
// isCaseSensitive (default):
//  - true
func ManyDefaultsPtr(
	str *string,
	splitsBy ...string,
) *strhelpercore.SplitResultOverview {
	return ManyPtr(
		str,
		&splitsBy,
		0,
		-1,
		true)
}

// Multiple split occur from the given array of splits.
//
// Basics of split("Hello World", " ") -> ["Hello", "World"] splitter will not be available in the result.
//
// limit :
//  - number of times split will performed for all
//  - if -1 then all split will occur
//
// splitStartsAt:
//  - where split searching will start from.
func ManyPtr(
	str *string,
	splitsBy *[]string,
	splitStartsAt,
	limit int,
	isCaseSensitive bool,
) *strhelpercore.SplitResultOverview {
	if isstr.EmptyPtr(str) || splitsBy == nil || len(*splitsBy) == 0 {
		return strhelpercore.NewEmptySplitResultOverview(str)
	}

	allIndexes := index.OfAllMany(
		str,
		splitsBy,
		splitStartsAt,
		limit,
		isCaseSensitive)

	if allIndexes == nil || !allIndexes.HasResult() {
		return strhelpercore.NewEmptySplitResultOverview(str)
	}

	indexesAsKeyMap := allIndexes.GetIndexesMapWhereIndexAsKey()
	possibleCapacity := len(*indexesAsKeyMap) + constants.ArbitraryCapacity5
	results := make([]*string, 0, possibleCapacity)
	nonResults := make([]*string, 0, possibleCapacity)
	splitResults := make([]*strhelpercore.SplitResult, 0, possibleCapacity)
	nonEmptySplitResults := make([]*strhelpercore.SplitResult, 0, possibleCapacity)
	lastIndexOfSplit := 0
	isLimitSet := limit > -1
	var isEmptyWord bool

	for ; splitStartsAt <= allIndexes.LastIndexFound; splitStartsAt++ {
		if isLimitSet && limit <= 0 {
			break
		}

		searchStr, isIndexExist := (*indexesAsKeyMap)[splitStartsAt]
		isIndexExist = isIndexExist && lastIndexOfSplit <= splitStartsAt
		if isIndexExist {
			word := (*str)[lastIndexOfSplit:splitStartsAt]
			isEmptyWord = word == "" || whitespace.IsAsciiWhitespaces(&word)

			splitResult := strhelpercore.SplitResult{
				SplitPrev: &word,
				Separator: &searchStr,
				Index:     splitStartsAt,
				IsEmpty:   isEmptyWord,
			}

			results = append(results, &word)
			splitResults = append(splitResults, &splitResult)

			if !isEmptyWord {
				nonResults = append(nonResults, &word)
				nonEmptySplitResults = append(nonEmptySplitResults, &splitResult)
			}

			lastIndexOfSplit = splitStartsAt + len(searchStr)
		}

		if isLimitSet && isIndexExist {
			limit--
		}
	}

	resultsOverview := strhelpercore.SplitResultOverview{
		Results:              &results,
		NonEmptyResults:      &nonResults,
		NonEmptySplitResults: &nonEmptySplitResults,
		SplitResults:         &splitResults,
		IsEmptyResult:        !allIndexes.HasResult(),
	}

	return &resultsOverview
}
