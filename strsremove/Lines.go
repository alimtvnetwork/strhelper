package strsremove

import (
	"gitlab.com/evatix-go/strhelper/ds/strhashset"
)

// Lines
//
// Creates new lines where removeStr will not appear.
//
// @count:
// - `-1` meaning remove all.
// - or else remove up to the given number only.
func Lines(
	lines []string,
	removeStr string,
	startsAt,
	count int,
	isCaseSensitive bool,
) []string {
	return ManyLines(
		lines,
		strhashset.NewUsingArray(removeStr),
		startsAt,
		count,
		isCaseSensitive)
}
