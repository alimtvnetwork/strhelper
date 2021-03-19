package splits

import (
	"gitlab.com/evatix-go/core"
	"gitlab.com/evatix-go/core/constants"

	"gitlab.com/evatix-go/strhelper/internal/indexinternal"
	"gitlab.com/evatix-go/strhelper/strto"
)

func LastByLimitPtr(
	s, separator *string,
	isCaseSensitive bool,
	limits int,
) *[]string {
	if limits == 0 {
		return core.EmptyStringsPtr()
	}

	if s == nil || *s == constants.EmptyString {
		return &[]string{constants.EmptyString}
	}

	if limits == 1 {
		return &[]string{*s}
	}

	isSepEmpty := *separator == constants.EmptyString

	if isSepEmpty {
		emptySeparatorResults := splitByLastUsingEmptySep(s, limits)

		return emptySeparatorResults
	}

	wholeTextLength := len(*s)
	searchTextLength := len(*separator)

	sendingContent := s
	sendingSearchTerm := separator

	if isCaseSensitive == false {
		// insensitive
		sendingContent = strto.LowerStrPtr(sendingContent)
		sendingSearchTerm = strto.LowerStrPtr(separator)
	}

	allFoundIndexes := indexinternal.OfLastAllCaseSensitiveUsingLengthPtr(
		sendingContent,
		sendingSearchTerm,
		constants.Zero,
		limits-1,
		wholeTextLength,
		searchTextLength)

	if allFoundIndexes == nil {
		return &[]string{*s}
	}

	length := len(*allFoundIndexes)

	if length == 0 {
		return &[]string{*s}
	}

	list := make([]string, length+1)
	splitIndex := 0
	incrementingIndex := 0
	foundIndex := 0
	for incrementingIndex, foundIndex = range *allFoundIndexes {
		// "[ab]found....1[ab]...found...2[ab]...found3
		// "...found3"
		// "...found...2"
		// "found....1"
		splitIndex = foundIndex + searchTextLength
		list[incrementingIndex] = (*s)[splitIndex:wholeTextLength]
		wholeTextLength = foundIndex
	}

	list[incrementingIndex+1] = (*s)[0:wholeTextLength]

	return &list
}
