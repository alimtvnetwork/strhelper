package strconcat

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
)

// Join concatenates the strings / elements of its first argument to a single string.
//
// @sep:
//  - used to strconcat each strings / elements.
//
// Copied from golang library (reference : https://bit.ly/3oPHGdy).
func JoinPtr(elements *[]string, sep string) string {
	if elements == nil {
		return constants.EmptyString
	}

	elementsLength := len(*elements)

	switch elementsLength {
	case 0:
		return constants.EmptyString
	case 1:
		return (*elements)[0]
	}

	n := len(sep) * (elementsLength - 1)
	elementsNonPtr := *elements

	for i := 0; i < elementsLength; i++ {
		n += len(elementsNonPtr[i])
	}

	var b strings.Builder
	b.Grow(n)
	b.WriteString(elementsNonPtr[0])
	for _, s := range elementsNonPtr[1:] {
		b.WriteString(sep)
		b.WriteString(s)
	}

	return b.String()
}
