package concat

import (
	"strings"

	"gitlab.com/evatix-go/strhelper/constants"
)

// Copied from golang library (reference : https://bit.ly/3oPHGdy).
// Join concatenates the elements of its first argument to create a single string. The separator
// string sep is placed between elements in the resulting string.
// Skip empty or whitespace string in elements.
func JoinStrPtrExceptEmpty(elements *[]*string, sep *string) string {
	elementsLength := len(*elements)

	switch elementsLength {
	case 0:
		return constants.EmptyString
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
		if s == nil || *s == constants.EmptyString || len(*s) == 0 {
			continue
		}

		b.WriteString(*sep)
		b.WriteString(*s)
	}

	return b.String()
}
