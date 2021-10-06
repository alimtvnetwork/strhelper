package strsremove

import (
	"gitlab.com/evatix-go/core/constants"

	"gitlab.com/evatix-go/strhelper/ds/strhashset"
)

func finalRemoveResultsCaseSensitive(
	lines []string,
	newLines []string,
	removeLinesHashset *strhashset.Hashset,
	startsAt int,
	count int,
	length int,
) []string {
	var line string

	isCountUnset := count == constants.InvalidNotFoundCase

	for i := startsAt; i < length; i++ {
		line = lines[i]

		if removeLinesHashset.Has(line) && (isCountUnset || count > 0) {
			count--
			continue
		}

		newLines = append(newLines, line)
	}

	return newLines
}
