package strremove

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
)

// SimpleManySplitsByNonEmptyTrimmed Remove as per
// removes then splits by the given separator
// and then returns trimmed output without empty strings.
func SimpleManySplitsByNonEmptyTrimmed(
	content string,
	splitsBy string,
	removeRequests ...string,
) []string {
	for _, remove := range removeRequests {
		content = strings.ReplaceAll(content, remove, constants.EmptyString)
	}

	return strings.Split(content, splitsBy)
}
