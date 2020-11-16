package isstr

import (
	"gitlab.com/evatix-go/strhelper/index"
	"gitlab.com/evatix-go/strhelper/strconst"
)

func Exists(
	s, findingString string,
	isCaseSensitive bool,
) bool {
	return index.Of(
		s,
		findingString,
		strconst.Zero,
		isCaseSensitive) > strconst.InvalidNotFoundCase
}
