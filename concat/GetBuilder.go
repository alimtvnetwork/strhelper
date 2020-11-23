package concat

import "strings"

// Get Builder with all the items added as (joined items with separator) + growth of additionalLength
func GetBuilder(elements *[]string, sep *string, additionalLength int) strings.Builder {
	elementsLength := len(*elements)

	n := len(*sep) * (elementsLength - 1)
	for i := 0; i < elementsLength; i++ {
		n += len((*elements)[i])
	}

	var b strings.Builder
	b.Grow(n + additionalLength)
	b.WriteString((*elements)[0])
	for _, s := range (*elements)[1:] {
		b.WriteString(*sep)
		b.WriteString(s)
	}

	return b
}
