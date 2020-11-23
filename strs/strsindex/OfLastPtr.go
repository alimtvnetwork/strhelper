package strsindex

import (
	"strings"

	"gitlab.com/evatix-go/strhelper/internal/pkg/isstrsinternal"
	"gitlab.com/evatix-go/strhelper/internal/pkg/panichelper"
	"gitlab.com/evatix-go/strhelper/strconst"
)

// OfLastPtr returns the last index where the string first found, doesn't care about the rest of the items once found.
func OfLastPtr(
	lines *[]string,
	searchTerm *string,
	lastIndexIncreasedBy int,
	isCaseSensitive bool,
) int {
	if isstrsinternal.Empty(lines) || searchTerm == nil {
		return strconst.InvalidNotFoundCase
	}

	length := len(*lines)

	if lastIndexIncreasedBy <= strconst.InvalidNotFoundCase || lastIndexIncreasedBy > length-1 {
		panichelper.LastIndexIncreasedByFailed(lastIndexIncreasedBy, length)
	}

	if !isCaseSensitive {
		// insensitive
		toLowerSearchTerm := strings.ToLower(*searchTerm)
		index := length - lastIndexIncreasedBy

		for ; index >= 0; index-- {
			if strings.ToLower((*lines)[index]) == toLowerSearchTerm {
				return index
			}
		}
	}

	index := length - lastIndexIncreasedBy

	for ; index >= 0; index-- {
		if (*lines)[index] == *searchTerm {
			return index
		}
	}

	return strconst.InvalidNotFoundCase
}
