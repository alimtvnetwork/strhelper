package strremove

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
)

func SimpleMany(content string, removeRequests ...string) string {
	for _, remove := range removeRequests {
		content = strings.ReplaceAll(content, remove, constants.EmptyString)
	}

	return content
}
