package concat

import (
	"strings"

	"gitlab.com/evatix-go/strhelper/strconst"
)

// Concatenates the strings / elements of to a single string using the @sep (separator).
//
// Skip @filterSkipMap elements.
//
// @sep:
//  - used to concat each strings / elements.
//
// Copied from golang library (reference : https://bit.ly/3oPHGdy).
func JoinPtrExceptFor(filterSkipMap *map[string]bool, elements *[]string, sep *string) string {
	elementsLength := len(*elements)
	if elementsLength == 0 {
		return strconst.EmptyString
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
