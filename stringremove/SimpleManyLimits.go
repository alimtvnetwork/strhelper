package strremove

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
)

func SimpleManyLimits(content string, limits int, removeRequests ...string) string {
	for _, remove := range removeRequests {
		content = strings.Replace(content, remove, constants.EmptyString, limits)
	}

	return content
}
