package concat

import (
	"strings"

	"gitlab.com/evatix-go/strhelper/strconst"
)

// Concatenates the strings / elements to a single string using @sep (separator).
//
// Skips nil or empty string in elements. (not the whitespace)
//
// @sep:
//  - used to concat each strings / elements.
//
// Copied from golang library (reference : https://bit.ly/3oPHGdy).
func JoinStrPtrExceptEmpty(elements *[]*string, sep *string) string {
	elementsLength := len(*elements)

	switch elementsLength {
	case 0:
		return strconst.EmptyString
	case 1:
		return *(*elements)[0]
	}

	n := len(*sep) * (elementsLength - 1)
	for i := 0; i < elementsLength; i++ {
		n += len(*(*elements)[i])
	}

	var b strings.Builder
	b.Grow(n)
	b.WriteString(*(*elements)[0])
	for _, s := range (*elements)[1:] {
		if s == nil || *s == strconst.EmptyString || len(*s) == 0 {
			continue
		}

		b.WriteString(*sep)
		b.WriteString(*s)
	}

	return b.String()
}
