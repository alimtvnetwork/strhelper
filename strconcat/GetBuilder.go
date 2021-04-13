package strconcat

import "strings"

// Get Builder with all the items added as (joined items with separator) + growth of additionalLength
func GetBuilder(
	elements *[]string,
	sep string,
	additionalLength int,
) strings.Builder {
	if elements == nil {
		var bEmpty strings.Builder
		bEmpty.Grow(additionalLength)

		return bEmpty
	}

	elementsNonPtr := *elements
	elementsLength := len(elementsNonPtr)
	n := len(sep) * (elementsLength - 1)
	for i := 0; i < elementsLength; i++ {
		n += len(elementsNonPtr[i])
	}

	var b strings.Builder
	b.Grow(n + additionalLength)
	b.WriteString(elementsNonPtr[0])
	for _, s := range elementsNonPtr[1:] {
		b.WriteString(sep)
		b.WriteString(s)
	}

	return b
}
