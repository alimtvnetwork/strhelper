package splits

import (
	"strings"

	"gitlab.com/evatix-go/core"
	"gitlab.com/evatix-go/core/constants"

	"gitlab.com/evatix-go/strhelper/internal/indexinternal"
)

func Last(s, separator string) *[]string {
	if s == "" {
		return core.EmptyStringsPtr()
	}
	return LastByLimitPtr(
		&s,
		&separator,
		constants.MinusOne)
}

func LastByLimitPtr(s, separator *string, limits int) *[]string {
	if s == nil || *s == "" || separator == nil || limits == 0 {
		return core.EmptyStringsPtr()
	}

	isSepEmpty := *separator == ""

	if isSepEmpty && limits == constants.MinusOne {
		emptySeparatorResults := strings.Split(*s, "")

		return &emptySeparatorResults
	}

	wholeTextLength := len(*s)
	searchTextLength := len(*separator)

	allFoundIndexes := indexinternal.
		OfLastAllCaseSensitiveUsingLengthPtr(
			s,
			separator,
			constants.Zero,
			limits,
			wholeTextLength,
			searchTextLength)

	if allFoundIndexes == nil {
		return core.EmptyStringsPtr()
	}

	list := make([]string, len(*allFoundIndexes)+1)
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
