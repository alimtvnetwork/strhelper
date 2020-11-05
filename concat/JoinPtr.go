package concat

import (
	"strings"

	"gitlab.com/evatix-go/strhelper/constants"
)

// Copied from golang library (reference : https://bit.ly/3oPHGdy).
// Join concatenates the elements of its first argument to create a single string. The separator
// string sep is placed between elements in the resulting string.
func JoinPtr(elements *[]string, sep *string) string {
	elementsLength := len(*elements)

	switch elementsLength {
	case 0:
		return constants.EmptyString
	case 1:
		return (*elements)[0]
	}

	n := len(*sep) * (elementsLength - 1)
	for i := 0; i < elementsLength; i++ {
		n += len((*elements)[i])
	}

	var b strings.Builder
	b.Grow(n)
	b.WriteString((*elements)[0])
	restOfTheElementsExceptFirst := (*elements)[1:]
	for _, s := range restOfTheElementsExceptFirst {
		b.WriteString(*sep)
		b.WriteString(s)
	}

	return b.String()
}
