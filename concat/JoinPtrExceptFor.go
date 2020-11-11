package concat

import (
	"strings"

	"gitlab.com/evatix-go/strhelper/constants"
)

// Copied from golang library (reference : https://bit.ly/3oPHGdy).
// Join concatenates the elements of its first argument to create a single string. The separator
// string sep is placed between elements in the resulting string.
// Skip given filterSkipMap items from elements
func JoinPtrExceptFor(filterSkipMap *map[string]bool, elements *[]string, sep *string) string {
	elementsLength := len(*elements)
	if elementsLength == 0 {
		return constants.EmptyString
	}

	n := len(*sep) * (elementsLength - 1)
	for i := 0; i < elementsLength; i++ {
		n += len((*elements)[i])
	}

	var b strings.Builder
	b.Grow(n)
	b.WriteString((*elements)[0])
	for _, s := range (*elements)[1:] {
		_, has := (*filterSkipMap)[s]
		if has {
			continue
		}

		b.WriteString(*sep)
		b.WriteString(s)
	}

	return b.String()
}
