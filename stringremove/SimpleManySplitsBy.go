package strremove

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
)

// SimpleManySplitsBy Remove as per removes then splits by the given separator
func SimpleManySplitsBy(
	content string,
	splitsBy string,
	removeRequests ...string,
) []string {
	for _, remove := range removeRequests {
		content = strings.ReplaceAll(content, remove, constants.EmptyString)
	}

	return strings.Split(content, splitsBy)
}
