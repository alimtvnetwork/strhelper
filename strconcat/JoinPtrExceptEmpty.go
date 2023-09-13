package strconcat

import (
	"strings"

	"gitlab.com/auk-go/core/constants"
)

// Concatenates the strings / elements of to a single string using the @sep (separator).
//
// Skip nil or empty string in elements.
//
// @sep:
//  - used to strconcat each strings / elements.
//
// Copied from golang library (reference : https://bit.ly/3oPHGdy).
func JoinPtrExceptEmpty(elements *[]string, sep string) string {
	elementsLength := len(*elements)

	switch elementsLength {
	case 0:
		return constants.EmptyString
	case 1:
		return (*elements)[0]
	}

	elementsNonPtr := *elements

	n := len(sep) * (elementsLength - 1)
	for i := 0; i < elementsLength; i++ {
		n += len(elementsNonPtr[i])
	}

	var b strings.Builder
	b.Grow(n)
	b.WriteString(elementsNonPtr[0])
	for _, s := range elementsNonPtr[1:] {
		if s == constants.EmptyString {
			continue
		}

		b.WriteString(sep)
		b.WriteString(s)
	}

	return b.String()
}
