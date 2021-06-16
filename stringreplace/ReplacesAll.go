package stringreplace

import (
	"strings"

	"gitlab.com/evatix-go/core/coredata/corestr"
)

func ReplacesAll(
	text string,
	replaceRequests ...corestr.KeyValuePair,
) string {
	if len(replaceRequests) == 0 {
		return text
	}

	for _, request := range replaceRequests {
		text = strings.ReplaceAll(text, request.Key, request.Value)
	}

	return text
}
