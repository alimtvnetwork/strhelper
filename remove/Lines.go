package remove

import (
	"gitlab.com/evatix-go/strhelper/ds/strhashset"
)

// Creates new lines where removeStr will not appear.
//
// @count:
// - `-1` meaning remove all.
// - or else remove up to the given number only.
func Lines(
	lines *[]string,
	removeStr *string,
	startsAt,
	count int,
	isCaseSensitive bool,
) *[]string {
	return ManyLines(
		lines,
		strhashset.NewWithValues(*removeStr),
		startsAt,
		count,
		isCaseSensitive)
}
