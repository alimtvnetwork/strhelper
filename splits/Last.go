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

	list := make([]string, len(*allFoundIndexes))
	splitIndex := 0

	for i, index := range *allFoundIndexes {
		// "[ab]found....1[ab]...found...2[ab]...found3
		// "...found3"
		// "...found...2"
		// "found....1"
		splitIndex = index + searchTextLength + 1
		word := (*s)[splitIndex:wholeTextLength]
		wholeTextLength -= index - 1

		list[i] = word
	}

	return &list
}
