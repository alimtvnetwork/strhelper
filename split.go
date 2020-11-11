package strhelper

import (
	"gitlab.com/evatix-go/strhelper/constants"
	"gitlab.com/evatix-go/strhelper/strhelpercore"
	"gitlab.com/evatix-go/strhelper/whitespace"
)

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
func MultipleQuickSplits(
	str string,
	splitsBy ...string,
) *strhelpercore.SplitResultOverview {
	return MultipleSplitsPtr(
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
func MultipleSplits(
	str string,
	splitStartsAt,
	limit int,
	isCaseSensitive bool,
	splitsBy ...string,
) *strhelpercore.SplitResultOverview {
	return MultipleSplitsPtr(
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
// limit :
//  - number of times split will performed for all
//  - if -1 then all split will occur
//
// splitStartsAt:
//  - where split searching will start from.
func MultipleQuickSplitsPtr(
	str *string,
	splitsBy ...string,
) *strhelpercore.SplitResultOverview {
	return MultipleSplitsPtr(
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
func MultipleSplitsPtr(
	str *string,
	splitsBy *[]string,
	splitStartsAt,
	limit int,
	isCaseSensitive bool,
) *strhelpercore.SplitResultOverview {
	if IsEmptyPtr(str) || splitsBy == nil || len(*splitsBy) == 0 {
		return strhelpercore.NewEmptySplitResultOverview(str)
	}

	allIndexes := MultiStrIndexesOfAllUsingSimpleArrayPtr(
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
			isEmptyWord = word == "" || whitespace.IsASCIIWhitespaceOnly(&word)

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
