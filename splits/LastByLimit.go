package splits

import (
	"strings"

	"gitlab.com/auk-go/core/constants"

	"gitlab.com/auk-go/strhelper/internal/indexinternal"
)

func LastByLimit(
	wholeText, separator string,
	isCaseSensitive bool,
	limits int,
) []string {
	if limits == 0 {
		return []string{}
	}

	if wholeText == constants.EmptyString {
		return []string{constants.EmptyString}
	}

	if limits == 1 {
		return []string{wholeText}
	}

	isSepEmpty := separator == constants.EmptyString

	if isSepEmpty {
		emptySeparatorResults := splitByLastUsingEmptySep(wholeText, limits)

		return emptySeparatorResults
	}

	wholeTextLength := len(wholeText)
	searchTextLength := len(separator)

	sendingContent := wholeText
	sendingSearchTerm := separator

	if isCaseSensitive == false {
		// insensitive
		sendingContent = strings.ToLower(sendingContent)
		sendingSearchTerm = strings.ToLower(separator)
	}

	allFoundIndexes := indexinternal.OfLastAllCaseSensitiveUsingLengthPtr(
		sendingContent,
		sendingSearchTerm,
		constants.Zero,
		limits-1,
		wholeTextLength,
		searchTextLength)

	length := len(allFoundIndexes)

	if length == 0 {
		return []string{wholeText}
	}

	// accessing direct without pointer increases performance
	list := make([]string, length+1)
	splitIndex := 0
	incrementingIndex := 0
	foundIndex := 0
	for incrementingIndex, foundIndex = range allFoundIndexes {
		// "[ab]found....1[ab]...found...2[ab]...found3
		// "...found3"
		// "...found...2"
		// "found....1"
		splitIndex = foundIndex + searchTextLength
		list[incrementingIndex] = wholeText[splitIndex:wholeTextLength]
		wholeTextLength = foundIndex
	}

	// the remaining ones will be at the end
	list[incrementingIndex+1] = wholeText[0:wholeTextLength]

	return list
}
