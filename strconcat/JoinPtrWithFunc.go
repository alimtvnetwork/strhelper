package strconcat

import (
	"strings"

	"gitlab.com/auk-go/core/constants"
)

// JoinPtrWithFunc Concatenates the strings / elements to
// a single string using a compiler function and @sep (separator).
//
// Where useful:
//  - for example, if we want to use trim function for all the items or probably some addition to the all the items then
//      this can be very useful.
//
// @sep:
//  - used to strconcat each strings / elements.
// @compiler:
//  - process each element to processed string.
//  - compiler function returns string, bool. Here the bool refers to isGoingToAdd if returns false then
//      item will be discarded from adding.
//
// Copied from golang library (reference : https://bit.ly/3oPHGdy).
func JoinPtrWithFunc(
	elements *[]string,
	sep string,
	compiler func(element string) (string, bool),
) string {
	if elements == nil {
		return constants.EmptyString
	}

	elementsLength := len(*elements)
	if elementsLength == 0 {
		return constants.EmptyString
	}

	n := len(sep) * (elementsLength - 1)
	for i := 0; i < elementsLength; i++ {
		n += len((*elements)[i])
	}

	var b strings.Builder
	b.Grow(n)
	b.WriteString((*elements)[0])
	for _, s := range (*elements)[1:] {
		result, isAdd := compiler(s)

		if !isAdd {
			continue
		}

		b.WriteString(sep)
		b.WriteString(result)
	}

	return b.String()
}
